import {  UsersServiceClient } from '../../proto/UsersServiceClientPb'
import { userLoginRequest } from '../../proto/users_pb'

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
    },

    async login(state: any, payload: any) {
        let client = new UsersServiceClient('http://0.0.0.0:5002', {}, {})
        let request = new userLoginRequest()
        request.setUsername('admin')
        request.setPassword('admin')
        let response = await client.login(request,{})
        console.log('response', response.getToken())
    }
}