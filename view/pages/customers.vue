<template>
  <div>
    <form ref="file-input-form">
      <input type="file" @input="upload_excel($event)" />
    </form>
    <button @click="export_excel">Export</button>
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
    this.getCustomers();
  },
  methods: {
    async getCustomers() {
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
    async upload_excel(event: any) {
      event.preventDefault();
      var reader = new FileReader();
      let dataSheet: any = [];

      reader.onload = async (e: any) => {
        var data = new Uint8Array(e.target!.result);

        var workbook = XLSX.read(data, { type: "array", cellDates: true });
        let sheetName = workbook.SheetNames[0];
        let worksheet = workbook.Sheets[sheetName];
        let customerCreateRequest = new CustomerCreateRequest();
        let customers: Array<Customer> = [];
        XLSX.utils.sheet_to_json(worksheet).map((customerData: any) => {
          let customer = new Customer();
          customer.setNik(customerData["NIK"]);
          customer.setNama(customerData["Nama Lengkap"]);
          customer.setNokk(customerData["No KK"] + "");
          customer.setStatus(customerData["Status"] + "");
          customer.setStatuskk(customerData["Status KK"]);
          customer.setTanggallahir(
            customerData["Tanggal Lahir"]
              ? new Date(customerData["Tanggal Lahir"]).toLocaleDateString()
              : ""
          );
          customer.setKotakab(customerData["Kab/Kota"]);
          customer.setKecamatan(customerData["Kecamatan"]);
          customer.setDesakelurahan(customerData["Desa/Kel"]);
          customer.setKampung(customerData["KP"]);
          customer.setRt(customerData["RT"] + "");
          customer.setRw(customerData["RW"] + "");
          customer.setKol(customerData["Kol"]);
          customer.setSyahidan(customerData["Syahidan"]);
          customer.setPj(customerData["PJ"]);
          console.log('customer', customer.toObject())
          customers.push(customer);
        });
        customerCreateRequest.setDataList(customers);
        this.$store.dispatch("addCustomer", customerCreateRequest);
      };

      reader.readAsArrayBuffer(event.target.files[0]);
      (this.$refs["file-input-form"] as any).reset();
    },

    async export_excel() {
      let payload = new CustomerGetRequest();
      payload.setPage(-1);
      payload.setPerPage(-1);
      let response = await (<Promise<CustomerGetResponse>>(
        this.$store.dispatch("getCustomers", payload)
      ));

      let customers = response.getDataList();

      let data: Array<Array<string>> = [
        [
          "ID",
          "NIK",
          "Nama Lengkap",
          "Status KK",
          "No KK",
          "Status",
          "Tanggal Lahir",
          "Kab/Kota",
          "Kecamatan",
          "Desa/Kel",
          "KP",
          "RT",
          "RW",
          "Kol",
          "Syahidan",
          "PJ"
        ],
      ];

      for (const customer of customers) {
        console.log('customer', customer.toObject())
        data.push([
          customer.getId(),
          customer.getNik(),
          customer.getNama(),
          customer.getStatuskk(),
          customer.getNokk(),
          customer.getStatus(),
          customer.getTanggallahir(),
          customer.getKotakab(),
          customer.getKecamatan(),
          customer.getDesakelurahan(),
          customer.getKampung(),
          customer.getRt(),
          customer.getRw(),
          customer.getKol(),
          customer.getSyahidan(),
          customer.getPj()
        ],)
      }
      var filename = "customers.xlsx";

      var wb = XLSX.utils.book_new();
      var ws_location = XLSX.utils.aoa_to_sheet(data);

      XLSX.utils.book_append_sheet(wb, ws_location, "Sheet1");
      XLSX.writeFile(wb, filename);
    },
  },
});
</script>