'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
    }
    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);
    }
    
    async submitTransaction() {
        const randomId1 = Math.floor(Math.random()*1000);
        const randomId2 = Math.floor(Math.random()*20000000);
        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'addModelCredibility',
            invokerIdentity: 'client0.org1.digiblocks.com',
            contractArguments: ["TCMAPSS_Round2", "alice", "13.5", "2.3", "0.5874392029285431"] ,
            readOnly: false
        };
        await this.sutAdapter.sendRequests(myArgs);
    }
    
    async cleanupWorkloadModule() {
        // NOOP
    }
}
function createWorkloadModule() {
    return new MyWorkload();
}
module.exports.createWorkloadModule = createWorkloadModule;