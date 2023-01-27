<template>
  <div>
    <form ref="file-input-form">
      <input type="file" @input="upload_excel($event)" />
    </form>
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
          <td>{{ customer.nama }}</td>
          <td>{{ customer.nokk }}</td>
          <td>
            {{ customer.kampung ? customer.kampung : customer.desakelurahan }}
          </td>
          <td>{{ customer.pj }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import {
  CustomerGetRequest,
  CustomerGetResponse,
  Customer,
  CustomerCreateRequest,
} from "../../proto/customers_pb";
import * as XLSX from "xlsx";

export default Vue.extend({
  name: "Customer",
  data: () => {
    type customerState = {
      page: number;
      totalPage: number;
      dataList: Array<Customer.AsObject>;
    };

    let state: customerState = {
      page: 0,
      totalPage: 0,
      dataList: [],
    };

    return state;
  },
  async mounted() {
    let payload = new CustomerGetRequest();
    payload.setPage(this.page);
    let response = await (<Promise<CustomerGetResponse>>(
      this.$store.dispatch("getCustomers", payload)
    ));

    if (response) {
      let obj = response.toObject();
      this.totalPage = obj.totalPage;
      obj.dataList.map((customer: Customer.AsObject) => {
        this.dataList.push(customer);
      });

      console.log("response", obj);
    }
  },
  methods: {
    getCustomers() {},
    async upload_excel(event: any) {
      event.preventDefault()
      var reader = new FileReader();
      let dataSheet: any = [];

      reader.onload = async (e: any) => {
        var data = new Uint8Array(e.target!.result);

        var workbook = XLSX.read(data, { type: "array", cellDates: true });
        let sheetName = workbook.SheetNames[0];
        let worksheet = workbook.Sheets[sheetName];
        let customerCreateRequest = new CustomerCreateRequest()
        let customers: Array<Customer> = []
        XLSX.utils.sheet_to_json(worksheet).map((customerData: any) => {
          let customer = new Customer()
          customer.setNo(customerData['No'])
          customer.setNik(customerData['NIK'])
          customer.setNama(customerData['Nama Lengkap'])
          customer.setNokk(customerData['No KK']+"")
          customer.setStatuskk(customerData['Status KK'])
          customer.setTanggallahir(customerData['Tanggal Lahir'] ? new Date(customerData['Tanggal Lahir']).toString(): "")
          customer.setUsia(customerData['Usia'])
          customer.setKotakab(customerData['Kab/ Kota'])
          customer.setKecamatan(customerData['Kecamatan'])
          customer.setDesakelurahan(customerData['Desa/ Kel'])
          customer.setKampung(customerData['KP'])
          customer.setRt(customerData['RT'])
          customer.setRw(customerData['RW'])
          customer.setKol(customerData['Kol'])
          customer.setSyahidan(customerData['Syahidan'])
          customer.setPj(customerData['PJ'])
          customers.push(customer)
        })
        customerCreateRequest.setDataList(customers)
        this.$store.dispatch("addCustomer", customerCreateRequest)
      };

      reader.readAsArrayBuffer(event.target.files[0]);
		  (this.$refs['file-input-form'] as any).reset()

    },
  },
});
</script>