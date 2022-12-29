var PROTO_PATH_USER = __dirname + '/../../proto/users.proto';
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
// Suggested options for similarity to existing grpc.load behavior
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH_USER,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);

console.log('protoDescriptor', protoDescriptor)

export const state = () => ({
    counter: 0
})

export const getters = {
    getCounter(state: any) {
        return state.counter
    }
}

export const mutations = {
    increment(state: any) {
        state.counter++
    }
}

export const actions = {
    async fetchCounter(state: any, payload: any) {
        // make request
        console.log('payload', payload)
        const res = { data: 10 };
        state.counter = res.data;
        return res.data;
    }
}