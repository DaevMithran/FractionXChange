import {getGasFee, MantrachainSdk} from '../helpers/sdk'
import {createDenomIfNotExists, genCoinDenom} from "../helpers/coinfactory";
import {getPairId} from '../helpers/liquidity';
import {updateAccountPrivileges, updateCoinRequiredPrivileges} from '../helpers/guard';
import {mintGuardSoulBondNft} from "../helpers/token";

/** OrderDirection enumerates order directions. */
export enum OrderDirection {
    /** ORDER_DIRECTION_UNSPECIFIED - ORDER_DIRECTION_UNSPECIFIED specifies unknown order direction */
    ORDER_DIRECTION_UNSPECIFIED = 0,
    /** ORDER_DIRECTION_BUY - ORDER_DIRECTION_BUY specifies buy(swap quote coin to base coin) order direction */
    ORDER_DIRECTION_BUY = 1,
    /** ORDER_DIRECTION_SELL - ORDER_DIRECTION_SELL specifies sell(swap base coin to quote coin) order direction */
    ORDER_DIRECTION_SELL = 2,
    UNRECOGNIZED = -1,
}


describe('Liquidity module', () => {
    let sdk: MantrachainSdk

    let baseCoinDenom = 'liquidity' + new Date().getTime().toString();
    // with this we manipulate the time and space
    // of this test environment according to our needs
    // to the infinity and beyond!
    let quoteCoinDenom = 'liquidity' + new Date().getTime().toString() + 1;

    let pairId = 0;

    let poolId;

    beforeAll(async () => {
        sdk = new MantrachainSdk()
        await sdk.init(process.env.API_URL, process.env.RPC_URL, process.env.WS_URL)

        await createDenomIfNotExists(sdk, sdk.clientAdmin, sdk.adminAddress, baseCoinDenom)
        await createDenomIfNotExists(sdk, sdk.clientAdmin, sdk.adminAddress, quoteCoinDenom)

        await sdk.clientAdmin.MantrachainCoinfactoryV1Beta1.tx.sendMsgMint({
            value: {
                sender: sdk.adminAddress,
                amount: {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                    amount: "10000000000000000000"
                }
            },
            fee: getGasFee()
        })

        await sdk.clientAdmin.MantrachainCoinfactoryV1Beta1.tx.sendMsgMint({
            value: {
                sender: sdk.adminAddress,
                amount: {
                    denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                    amount: "10000000000000000000"
                }
            },
            fee: getGasFee()
        })

        await updateCoinRequiredPrivileges(
            sdk,
            sdk.clientAdmin,
            sdk.adminAddress,
            genCoinDenom(sdk.adminAddress, baseCoinDenom),
            [1, 1]
        )
        await updateCoinRequiredPrivileges(
            sdk,
            sdk.clientAdmin,
            sdk.adminAddress,
            genCoinDenom(sdk.adminAddress, quoteCoinDenom),
            [1, 1]
        )

        await updateAccountPrivileges(sdk, sdk.clientAdmin, sdk.adminAddress, sdk.recipientAddress, [1, 1])

        await sdk.clientAdmin.CosmosBankV1Beta1.tx.sendMsgSend({
            value: {
                fromAddress: sdk.adminAddress,
                toAddress: sdk.recipientAddress,
                amount: [
                    {
                        denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                        amount: '100000000000'
                    },
                    {
                        denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        amount: '100000000000'
                    }
                ]
            },
            fee: getGasFee()
        })
    })

    describe('Admin', () => {
        test('should be able to create pair for existing denoms', async () => {

            await sdk.clientAdmin.MantrachainLiquidityV1Beta1.tx.sendMsgCreatePair({
                value: {
                    creator: sdk.adminAddress,
                    baseCoinDenom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                    quoteCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom)
                },
                fee: getGasFee()
            })

            pairId = await getPairId(
                sdk.clientAdmin,
                genCoinDenom(sdk.adminAddress, baseCoinDenom),
                genCoinDenom(sdk.adminAddress, quoteCoinDenom)
            )

            expect(pairId).toBeGreaterThan(0);
        });

        test('should throw when trying to create already existing pair for existing denoms', async () => {

            const res = await sdk.clientAdmin.MantrachainLiquidityV1Beta1.tx.sendMsgCreatePair({
                value: {
                    creator: sdk.adminAddress,
                    baseCoinDenom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                    quoteCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom)
                },
                fee: getGasFee()
            })

            expect(res.code).not.toBe(0)
            expect(res.rawLog).toMatch(/pair already exists/);
        });

        // TODO currently this will always fail as it is possible to create pair for non existing denoms
        // test('should get error when trying to create pair for non existing denoms', async () => {
        //     const res = sdk.clientAdmin.MantrachainLiquidityV1Beta1.tx.sendMsgCreatePair({
        //         value: {
        //             creator: sdk.adminAddress,
        //             baseCoinDenom: genCoinDenom(sdk.adminAddress, 'asdasdasd'),
        //             quoteCoinDenom: genCoinDenom(sdk.adminAddress, 'asdasdasdasd')
        //         }
        //     })
        //
        //     await expect(res).rejects.toThrow()
        // });

        test('should be able to create pool for existing pair', async () => {

            let z = await sdk.clientAdmin.MantrachainLiquidityV1Beta1.tx.sendMsgCreatePool(
                {
                    value: {
                        creator: sdk.adminAddress,
                        pairId: pairId,
                        depositCoins: [
                            {
                                denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                                amount: "1000000000"
                            },
                            {
                                denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                                amount: "1000000000"
                            }
                        ]
                    },
                    fee: getGasFee()
                }
            )

            const resp = await sdk.clientAdmin.MantrachainLiquidityV1Beta1.query.queryPools({
                pair_id: pairId.toString()
            });

            const poolIndex = resp.data.pools.length - 1;
            poolId = resp.data.pools[poolIndex].id;

            expect(resp.data.pools.find(pool => pool.creator == sdk.adminAddress)).toBeTruthy()
        })

        test('should be able to deposit liquidity to existing pool', async () => {
            const balanceOfBaseCoinsBefore = await sdk.clientAdmin.CosmosBankV1Beta1.query.queryBalance(
                sdk.adminAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            await sdk.clientAdmin.MantrachainLiquidityV1Beta1.tx.sendMsgDeposit(
                {
                    value: {
                        depositor: sdk.adminAddress,
                        poolId: poolId,
                        depositCoins: [
                            {
                                denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                                amount: "1000000000"
                            },
                            {
                                denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                                amount: "1000000000"
                            }
                        ]
                    },
                    fee: getGasFee()
                }
            )

            const balanceOfBaseCoinsAfter = await sdk.clientAdmin.CosmosBankV1Beta1.query.queryBalance(
                sdk.adminAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )
            expect(Number(balanceOfBaseCoinsBefore.data.balance.amount)).toBeGreaterThan(Number(balanceOfBaseCoinsAfter.data.balance.amount))
        })

        test('should be able to withdraw liquidity from existing pool', async () => {
            const resp = await sdk.clientAdmin.MantrachainLiquidityV1Beta1.query.queryPool(poolId);
            const pool = resp.data.pool;

            const balanceOfBaseCoinsBefore = await sdk.clientAdmin.CosmosBankV1Beta1.query.queryBalance(
                sdk.adminAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            await sdk.clientAdmin.MantrachainLiquidityV1Beta1.tx.sendMsgWithdraw(
                {
                    value: {
                        withdrawer: sdk.adminAddress,
                        poolId: poolId,
                        poolCoin: {
                            denom: pool.pool_coin_denom,
                            amount: "5000000"
                        }
                    },
                    fee: getGasFee()
                }
            )

            const balanceOfBaseCoinsAfter = await sdk.clientAdmin.CosmosBankV1Beta1.query.queryBalance(
                sdk.adminAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )
            expect(Number(balanceOfBaseCoinsBefore.data.balance.amount)).toBeLessThan(Number(balanceOfBaseCoinsAfter.data.balance.amount))
        })

        test('should be able to create rangedPool for existing pair', async () => {
            await sdk.clientAdmin.MantrachainLiquidityV1Beta1.tx.sendMsgCreateRangedPool(
                {
                    value: {
                        creator: sdk.adminAddress,
                        pairId: pairId,
                        depositCoins: [
                            {
                                denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                                amount: "1000000000"
                            },
                            {
                                denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                                amount: "1000000000"
                            }
                        ],
                        minPrice: '1400000000000000000',
                        maxPrice: '1600000000000000000',
                        initialPrice: '1500000000000000000',
                    },
                    fee: getGasFee()
                }
            )

            const resp = await sdk.clientAdmin.MantrachainLiquidityV1Beta1.query.queryPools({
                pair_id: pairId.toString()
            });

            expect(resp.data.pools.find(pool => pool.creator == sdk.adminAddress)).toBeTruthy()
        })
    })

    describe('User', () => {
        test('should throw when trying to create pairs', async () => {
            const res = await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgCreatePair({
                value: {
                    creator: sdk.recipientAddress,
                    baseCoinDenom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                    quoteCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom)
                },
                fee: getGasFee()
            })

            expect(res.code).not.toBe(0)
            // expect(res.rawLog).toMatch(/unauthorized/);
        })

        test('should throw when trying to create pools', async () => {

            const res = await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgCreatePool(
                {
                    value: {
                        creator: sdk.recipientAddress,
                        pairId: pairId,
                        depositCoins: [
                            {
                                denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                                amount: "100000000"
                            },
                            {
                                denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                                amount: "100000000"
                            }
                        ]
                    },
                    fee: getGasFee()
                }
            )

            expect(res.code).not.toBe(0)
        })

        test('should throw when trying to create rangedPools', async () => {
            const res = await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgCreateRangedPool(
                {
                    value: {
                        creator: sdk.recipientAddress,
                        pairId: pairId,
                        depositCoins: [
                            {
                                denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                                amount: "1000000"
                            },
                            {
                                denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                                amount: "1000000"
                            }
                        ],
                        minPrice: '1400000000000000000',
                        maxPrice: '1600000000000000000',
                        initialPrice: '1500000000000000000',
                    },
                    fee: getGasFee()
                }
            )

            expect(res.code).not.toBe(0)
            // expect(res.rawLog).toMatch(/unauthorized/);
        })

        test('should be able to deposit liquidity to existing pool', async () => {

            const balanceOfBaseCoinsBefore = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgDeposit(
                {
                    value: {
                        depositor: sdk.recipientAddress,
                        poolId: poolId,
                        depositCoins: [
                            {
                                denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                                amount: "1000000"
                            },
                            {
                                denom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                                amount: "1000000"
                            }
                        ]
                    },
                    fee: getGasFee()
                }
            )

            const balanceOfBaseCoinsAfter = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )
            expect(Number(balanceOfBaseCoinsBefore.data.balance.amount)).toBeGreaterThan(Number(balanceOfBaseCoinsAfter.data.balance.amount))
        })

        test('should be able to withdraw liquidity from existing pool after deposit', async () => {
            const resp = await sdk.clientAdmin.MantrachainLiquidityV1Beta1.query.queryPool(poolId);

            const pool = resp.data.pool;

            const balanceOfBaseCoinsBefore = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: pool.pool_coin_denom
                }
            )

            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgWithdraw(
                {
                    value: {
                        withdrawer: sdk.recipientAddress,
                        poolId: poolId,
                        poolCoin: {
                            denom: pool.pool_coin_denom,
                            amount: "200000000"
                        }
                    },
                    fee: getGasFee()
                }
            )

            const balanceOfBaseCoinsAfter = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: pool.pool_coin_denom
                }
            )

            expect(Number(balanceOfBaseCoinsBefore.data.balance.amount)).toBeGreaterThan(Number(balanceOfBaseCoinsAfter.data.balance.amount))
        })

        test('should be able to make limit order', async () => {
            const balanceOfBaseCoinsBefore = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgLimitOrder({
                    value: {
                        orderer: sdk.recipientAddress,
                        pairId: pairId,
                        direction: OrderDirection.ORDER_DIRECTION_SELL,
                        offerCoin: {
                            denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                            amount: "1000000"
                        },
                        demandCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        price: '1400000000000000000',
                        amount: '1000000',
                        orderLifespan: undefined
                    },
                    fee: getGasFee()
                }
            )

            const balanceOfBaseCoinsAfter = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            expect(Number(balanceOfBaseCoinsBefore.data.balance.amount)).toBeGreaterThan(Number(balanceOfBaseCoinsAfter.data.balance.amount))
        })

        test('should be able to make market order', async () => {
            const balanceOfBaseCoinsBefore = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            const z = await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgMarketOrder(
                {
                    value: {
                        orderer: sdk.recipientAddress,
                        pairId: pairId,
                        direction: OrderDirection.ORDER_DIRECTION_SELL,
                        offerCoin: {
                            denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                            amount: "1000000"
                        },
                        demandCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        amount: '1000000',
                        orderLifespan: undefined
                    },
                    fee: getGasFee()
                }
            )

            const balanceOfBaseCoinsAfter = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            expect(Number(balanceOfBaseCoinsBefore.data.balance.amount)).toBeGreaterThan(Number(balanceOfBaseCoinsAfter.data.balance.amount))
        })

        test('should be able to make MM order', async () => {
            const balanceOfBaseCoinsBefore = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgMMOrder(
                {
                    value: {
                        orderer: sdk.recipientAddress,
                        pairId: pairId,
                        direction: OrderDirection.ORDER_DIRECTION_SELL,
                        offerCoin: {
                            denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                            amount: "1000000"
                        },
                        demandCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        price: '1400000000000000000',
                        amount: '1000000',
                        orderLifespan: {}
                    },
                    fee: getGasFee()
                }
            )

            const balanceOfBaseCoinsAfter = await sdk.clientRecipient.CosmosBankV1Beta1.query.queryBalance(
                sdk.recipientAddress, {
                    denom: genCoinDenom(sdk.adminAddress, baseCoinDenom)
                }
            )

            expect(Number(balanceOfBaseCoinsBefore.data.balance.amount)).toBeGreaterThan(Number(balanceOfBaseCoinsAfter.data.balance.amount))
        })

        test('should be able to cancel order', async () => {
            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgMMOrder(
                {
                    value: {
                        orderer: sdk.recipientAddress,
                        pairId: pairId,
                        direction: OrderDirection.ORDER_DIRECTION_SELL,
                        offerCoin: {
                            denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                            amount: "1000000"
                        },
                        demandCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        price: '1500000000000000000',
                        amount: '1000000',
                        orderLifespan: {seconds: 120}
                    },
                    fee: getGasFee()
                }
            )

            const orders = await sdk.clientRecipient.MantrachainLiquidityV1Beta1.query.queryOrders(pairId.toString())

            const order = orders.data.orders.find(o => o.type == 'ORDER_TYPE_MM')

            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgCancelOrder({
                value: {
                    orderer: sdk.recipientAddress,
                    pairId: pairId,
                    orderId: order.id
                },
                fee: getGasFee()
            })

            await new Promise((r) => setTimeout(r, 7000));

            const ordersAfter = await sdk.clientRecipient.MantrachainLiquidityV1Beta1.query.queryOrders(pairId.toString())
            expect(ordersAfter.data.orders.length).toBe(0)
        })

        test('should be able to cancel all orders', async () => {
            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgLimitOrder(
                {
                    value: {
                        orderer: sdk.recipientAddress,
                        pairId: pairId,
                        direction: OrderDirection.ORDER_DIRECTION_SELL,
                        offerCoin: {
                            denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                            amount: "1000000"
                        },
                        demandCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        price: '1500000000000000000',
                        amount: '1000000',
                        orderLifespan: {seconds: 120}
                    },
                    fee: getGasFee()
                }
            )

            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgMarketOrder(
                {
                    value: {
                        orderer: sdk.recipientAddress,
                        pairId: pairId,
                        direction: OrderDirection.ORDER_DIRECTION_SELL,
                        offerCoin: {
                            denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                            amount: "1000000"
                        },
                        demandCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        amount: '1000000',
                        orderLifespan: {seconds: 120}
                    },
                    fee: getGasFee()
                }
            )

            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgMMOrder(
                {
                    value: {
                        orderer: sdk.recipientAddress,
                        pairId: pairId,
                        direction: OrderDirection.ORDER_DIRECTION_SELL,
                        offerCoin: {
                            denom: genCoinDenom(sdk.adminAddress, baseCoinDenom),
                            amount: "1000000"
                        },
                        demandCoinDenom: genCoinDenom(sdk.adminAddress, quoteCoinDenom),
                        price: '1500000000000000000',
                        amount: '1000000',
                        orderLifespan: {seconds: 120}
                    },
                    fee: getGasFee()
                }
            )


            await sdk.clientRecipient.MantrachainLiquidityV1Beta1.tx.sendMsgCancelAllOrders({
                value: {
                    orderer: sdk.recipientAddress
                },
                fee: getGasFee()
            })

            await new Promise((r) => setTimeout(r, 7000));

            const ordersAfter = await sdk.clientRecipient.MantrachainLiquidityV1Beta1.query.queryOrders(pairId.toString())

            expect(ordersAfter.data.orders.length).toBe(0)
        })
    })
})