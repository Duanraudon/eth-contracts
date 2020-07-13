// const assertRevert = require('./helpers/assertRevert');
// const EthCrossChainManager = artifacts.require('./../contracts/core/v1.0/CrossChainManager/EthCrossChainManager');

// contract('EthCrossChain', (accounts) => {

//   before(async function () {
//     this.crossChain = await EthCrossChainManager.new({from:accounts[0], value: web3.utils.toWei('0', 'ether'), gas:60000000, gasPrice:50});
//   });

//   describe('InitGenesisBlock', function () {

//     const genesisHeader = '0x010000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c8365b000000001dac2b7c00000000fd17057b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a372c2263223a322c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a372c226964223a2231323035303364313031383338383037656334303739613436666539386436626439613036393061626362643863653136653066626334353230633763376566373838356462227d2c7b22696e646578223a312c226964223a2231323035303363366536383165353135346566626136346337356230616131636135343438396261653736353330373764313664646439373236663336356265333036323264227d2c7b22696e646578223a362c226964223a2231323035303361366231623065316336393737663434663336323332613566336236316236613835396234636535313437633439616363666139613432663438336631323034227d2c7b22696e646578223a342c226964223a2231323035303334343031376363636138323064393066306562623436316466343633333762303932336230616532626365353833636565316132363234633932303865323038227d2c7b22696e646578223a322c226964223a2231323035303333336334343833376464623934616435666130656234363062306634393135346639303530333631396434643263386565303833333066623831353834316432227d2c7b22696e646578223a332c226964223a2231323035303331326631303233393531333134336330323938346263346561396438353438383366636466343937333264633732376466613734373438326663383037653634227d2c7b22696e646578223a352c226964223a2231323035303266663764666337303562623561633638643265383932333063363632393939616562313832383431333165396663653934656639666166356239393137353364227d5d2c22706f735f7461626c65223a5b322c362c342c352c342c352c372c372c322c322c322c312c332c322c312c352c352c342c322c342c352c312c332c312c362c372c362c332c372c332c342c312c372c342c312c362c372c342c352c352c342c362c372c362c332c312c342c322c352c362c352c352c332c312c342c352c362c332c362c372c352c322c342c332c372c372c362c342c312c372c332c342c372c362c342c332c362c362c332c322c332c312c352c352c322c372c332c332c312c312c362c322c372c312c362c322c352c342c312c322c322c332c322c312c375d2c224d6178426c6f636b4368616e676556696577223a33303030307d7dfd9e5473b163f591a8829d83288809d97c20ab2a'
//     const genesisPubKey = '0x12050412f10239513143c02984bc4ea9d854883fcdf49732dc727dfa747482fc807e64ea4890067da5d75e3ce914dea1ab10d5d0774fd85e5a897c7cedb7a0ad694ef512050433c44837ddb94ad5fa0eb460b0f49154f90503619d4d2c8ee08330fb815841d2176c8ae35f754fa8876a25f63a25bcfa37548e2291f7ee55ddb790a36a88625512050444017ccca820d90f0ebb461df46337b0923b0ae2bce583cee1a2624c9208e208925fc299c1ce5437244e6774a34b2d51e45607eb9063606b943d015c8e391091120504a6b1b0e1c6977f44f36232a5f3b61b6a859b4ce5147c49accfa9a42f483f1204a349865a96ce5e181b98505ee8d64c8c87062cb3e67eccd2ad26c758b3b62cc7120504c6e681e5154efba64c75b0aa1ca54489bae7653077d16ddd9726f365be30622d3dc5c23474427b87e704319a440669092cf994f81acd8a567c96ddaaa8313349120504d101838807ec4079a46fe98d6bd9a0690abcbd8ce16e0fbc4520c7c7ef7885db612e135ec789514deb3a324d119519a3bc98dcb2607854e12d2b714bfe5b1c13120504ff7dfc705bb5ac68d2e89230c662999aeb18284131e9fce94ef9faf5b991753d39d4cfd17858edf9cfc1cf755dcb0162fceab314a028e8bc35e1c2b42fda8dc4'

//     it('InitGenesisBlock correctly', async function () {
//       const result = await this.crossChain.InitGenesisBlock(genesisHeader, genesisPubKey);
//     });

//     it('InitGenesisBlock throws error on InitGenesisBlock twice', async function () {
//       await assertRevert(this.crossChain.InitGenesisBlock(genesisHeader, genesisPubKey));
//     });

//   });

//    describe('SyncBlockHeader', function () {

//     const blockHeader = '0x0100000004000000000000001c87ef52276fee3dedb7d9108613a8340ff1eb5e02055f0ec8c8245ac1fdc5140000000000000000000000000000000000000000000000000000000000000000e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855aa06222b544e6b74883e5c3edcfcf996ec7bcc39b46fe236d1586dac2c32ea2061bab65d020000004f944d35d29cb94cfd0c017b226c6561646572223a372c227672665f76616c7565223a22424c4c74563248565a414f4b454c5842685471334654496e73626c7330706253333545787059506e43535a746c49777a657773673153354c6269754c34786b77542b623143496751764b532b3672457572506653585a633d222c227672665f70726f6f66223a2266724c71386a4b68425855535443554547374f376b4e3479736962307a3979477a6a6f757032664c327a53344474316b654c667179352b49414e6b5a4d704c752f634c37334d7034705268774f3170633551463038773d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d0000000000000000000000000000000000000000'
//     const signatures = '0xcfc4a5a2b0226bcf9002421ffa1d4b4c927202d733fcf307345bbc567726f11f1fd84f8f98858271c72575a9f1ada173d7a1f7e5649c37cf5e41c4e63b75bfcc01a3eeca37f1522eca3e8b1b015e108d650fd887fea46b5227d8f529a1a488821a79084e0cd5b82bc6232b9998e3900e9802de26efbc1db483053d2b37f43c852801fb3060f0a6fd4db8e4cfd17dcb651b094d6632ff34e1b4da6d0dc56c437148275686bb841536a75b881416a260fa0f7853efe5ada219ae61365b53d3cd5cf7280192ac26b5c6afff7b93e1ebbb51b3c18645678d7a33c502749669eb35f218c7377bc3b21050ec264fa043004de4a38e9a5f5dc43db56b2781d05af6ccc1bcc0c8011890c5b673940ef5434de714b6458cefa43deb8f831251f0241ebc4d88347f9e61549f2a760161c8982733cc3ca701884c314fb602a790c9e60a86ec09b0e07201'

//     it('SyncBlockHeader correctly', async function () {
//       const result = await this.crossChain.SyncBlockHeader(blockHeader, signatures);
//     });

//     it('SyncBlockHeader throws error on sync same block height header twice', async function () {
//       await assertRevert(this.crossChain.SyncBlockHeader(blockHeader, signatures));
//     });

//   });

//  describe('SyncAndVerify', function () {

//    const blockHeader = '0x0100000004000000000000001c87ef52276fee3dedb7d9108613a8340ff1eb5e02055f0ec8c8245ac1fdc5140000000000000000000000000000000000000000000000000000000000000000e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855aa06222b544e6b74883e5c3edcfcf996ec7bcc39b46fe236d1586dac2c32ea2061bab65d020000004f944d35d29cb94cfd0c017b226c6561646572223a372c227672665f76616c7565223a22424c4c74563248565a414f4b454c5842685471334654496e73626c7330706253333545787059506e43535a746c49777a657773673153354c6269754c34786b77542b623143496751764b532b3672457572506653585a633d222c227672665f70726f6f66223a2266724c71386a4b68425855535443554547374f376b4e3479736962307a3979477a6a6f757032664c327a53344474316b654c667179352b49414e6b5a4d704c752f634c37334d7034705268774f3170633551463038773d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d0000000000000000000000000000000000000000'
//    const signatures = '0xcfc4a5a2b0226bcf9002421ffa1d4b4c927202d733fcf307345bbc567726f11f1fd84f8f98858271c72575a9f1ada173d7a1f7e5649c37cf5e41c4e63b75bfcc01a3eeca37f1522eca3e8b1b015e108d650fd887fea46b5227d8f529a1a488821a79084e0cd5b82bc6232b9998e3900e9802de26efbc1db483053d2b37f43c852801fb3060f0a6fd4db8e4cfd17dcb651b094d6632ff34e1b4da6d0dc56c437148275686bb841536a75b881416a260fa0f7853efe5ada219ae61365b53d3cd5cf7280192ac26b5c6afff7b93e1ebbb51b3c18645678d7a33c502749669eb35f218c7377bc3b21050ec264fa043004de4a38e9a5f5dc43db56b2781d05af6ccc1bcc0c8011890c5b673940ef5434de714b6458cefa43deb8f831251f0241ebc4d88347f9e61549f2a760161c8982733cc3ca701884c314fb602a790c9e60a86ec09b0e07201'
//    const proof = '0x';
//    const position = 0;
//    const merkelValues = '0x';
//    const blockHeight = 0;

//    it('SyncBlockHeader correctly', async function () {
//      const result = await this.crossChain.SyncAndVerify(blockHeader, signatures, proof, position, merkelValues, blockHeight);
//    });

//    it('SyncAndVerify throws error on sync same block height header twice', async function () {
//      await assertRevert(this.crossChain.SyncAndVerify(blockHeader, signatures));
//    });

//  });
// });