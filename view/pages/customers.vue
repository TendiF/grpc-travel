<template>
  <div>
    <table>
        <thead>
        <tr>
            <th>Nama</th>
            <th>KK</th>
            <th>Desa/Kampung</th>
            <th>PJ</th>
        </tr>
        </thead>
        <tbody>
          <tr v-for="customer in dataList" v-bind:key="customer._id">
              <td>{{customer.nama}}</td>
              <td>{{customer.nokk}}</td>
              <td>{{customer.kampung ? customer.kampung : customer.desakelurahan }}</td>
              <td>{{customer.pj}}</td>
          </tr>
        </tbody>
    </table>

  </div>
</template>


<script lang="ts">
import Vue from 'vue'
import { CustomerGetRequest, CustomerGetResponse, CustomerCreateRequest } from '../../proto/customers_pb'


export default Vue.extend({
  name: 'Customer',
  data: () => {
    type customerState = {
      page: number,
      totalPage: number,
      dataList:  Array<CustomerCreateRequest.AsObject>
    }

    let state: customerState = {
      page: 0,
      totalPage: 0,
      dataList: []
    }

    return state
  },
  async mounted(){
    let payload = new CustomerGetRequest()
    payload.setPage(this.page)
    let response = await <Promise<CustomerGetResponse>>this.$store.dispatch('getCustomers', payload)
    
    if(response){
      let obj = response.toObject()
      this.totalPage = obj.totalPage
      obj.dataList.map((customer: CustomerCreateRequest.AsObject) => {
        this.dataList.push(customer)
      })
      
      console.log('response', obj)
    }
  },
  methods: {
    getCustomers(){

    }
  }
})
</script>