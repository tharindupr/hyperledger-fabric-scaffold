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
            contractFunction: 'createLocalModel',
            invokerIdentity: 'client0.org1.digiblocks.com',
            contractArguments: [`${this.workerIndex}___${randomId2}`, "3","3524", "alice", "80"] ,
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