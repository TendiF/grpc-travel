import {  CustomersServiceClient } from '../../proto/CustomersServiceClientPb'
import {  UsersServiceClient } from '../../proto/UsersServiceClientPb'
import { userLoginRequest } from '../../proto/users_pb'
import { CustomerGetRequest, CustomerCreateRequest } from '../../proto/customers_pb'

const baseUrl = 'http://0.0.0.0:5002'

type IndexStoreState = {
    token: string
}

let stateObj: IndexStoreState = {
    token: '',
}

type ActionContext = {
    commit: Function,
    dispatch: Function,
    getters: any,
    state: IndexStoreState,
}
export const state = () => (stateObj)

export const getters = {
    getToken(state: IndexStoreState) {
        return state.token
    }
}

export const mutations = {
    setToken(state: IndexStoreState, token: string) {
        state.token = token
    }
}

export const actions = {
    async login(context: ActionContext, payload: userLoginRequest) {
        let client = new UsersServiceClient(baseUrl, {}, {})
        return client.login(payload,{}).catch(error => {
            console.log('need to handle', error)
        })
    },

    async getCustomers(context: ActionContext, payload: CustomerGetRequest) {
        let client = new CustomersServiceClient(baseUrl, {}, {})
        return await client.get(payload, {
            authorization: context.getters.getToken
        }).catch(error => {
            if(error.code && error.code === 16){
                (this as any).$router.push('/')
            }
            console.log('need to handle', error)
        })
    },

    async addCustomer(context: ActionContext, payload: CustomerCreateRequest) {
        let client = new CustomersServiceClient(baseUrl, {}, {})

        return await client.create(payload, {
            authorization: context.getters.getToken
        }).catch(error => {
            if(error.code && error.code === 16){
                (this as any).$router.push('/')
            }
            console.log('need to handle', error)
        })

    }


}