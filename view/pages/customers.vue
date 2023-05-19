<template>
  <div>
    <form ref="file-input-form">
      <input type="file" @input="upload_excel($event)" />
    </form>
    <button @click="export_excel">Export</button>
    <div class="flex flex-col">
      <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
          <div class="overflow-hidden">
            <table class="min-w-full text-left text-sm font-light">
              <thead class="border-b font-medium dark:border-neutral-500">
                <tr>
                  <th scope="col" class="px-6 py-4">Nama</th>
                  <th scope="col" class="px-6 py-4">KK</th>
                  <th scope="col" class="px-6 py-4">Desa/Kampung</th>
                  <th scope="col" class="px-6 py-4">PJ</th>
                  <th scope="col" class="px-6 py-4">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  class="border-b dark:border-neutral-500"
                  v-for="customer in dataList"
                  v-bind:key="customer._id"
                >
                  <td class="whitespace-nowrap px-6 py-4 font-medium">
                    {{ customer.nama }}
                  </td>
                  <td class="whitespace-nowrap px-6 py-4">
                    {{ customer.nokk }}
                  </td>
                  <td class="whitespace-nowrap px-6 py-4">
                    {{
                      customer.kampung
                        ? customer.kampung
                        : customer.desakelurahan
                    }}
                  </td>
                  <td class="whitespace-nowrap px-6 py-4">{{ customer.pj }}</td>
                  <td class="whitespace-nowrap px-6 py-4">
                    <button
                      class="bg-blue-500 text-sm active:bg-gray-700 cursor-pointer font-regular text-white px-4 py-2 rounded uppercase"
                    >
                      TUNAI
                    </button>
                    <button
                      class="bg-green-500 text-sm active:bg-gray-700 cursor-pointer font-regular text-white px-4 py-2 rounded uppercase"
                    >
                      TERBINA
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
    <div
      class="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6"
    >
      <div class="flex flex-1 justify-between sm:hidden">
        <a
          href="#"
          class="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
          >Previous</a
        >
        <a
          href="#"
          class="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
          >Next</a
        >
      </div>
      <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            Showing
            <span class="font-medium">1</span>
            to
            <span class="font-medium">10</span>
            of
            <span class="font-medium">97</span>
            results
          </p>
        </div>
        <div>
          <nav
            class="isolate inline-flex -space-x-px rounded-md shadow-sm"
            aria-label="Pagination"
          >
            <a
              href="#"
              class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
            >
              <span class="sr-only">Previous</span>
              <svg
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
              >
                <path
                  fill-rule="evenodd"
                  d="M12.79 5.23a.75.75 0 01-.02 1.06L8.832 10l3.938 3.71a.75.75 0 11-1.04 1.08l-4.5-4.25a.75.75 0 010-1.08l4.5-4.25a.75.75 0 011.06.02z"
                  clip-rule="evenodd"
                />
              </svg>
            </a>
            <!-- Current: "z-10 bg-indigo-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600", Default: "text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0" -->
            <a
              href="#"
              aria-current="page"
              class="relative z-10 inline-flex items-center bg-indigo-600 px-4 py-2 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >1</a
            >
            <a
              href="#"
              class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
              >2</a
            >
            <a
              href="#"
              class="relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 md:inline-flex"
              >3</a
            >
            <span
              class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-700 ring-1 ring-inset ring-gray-300 focus:outline-offset-0"
              >...</span
            >
            <a
              href="#"
              class="relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 md:inline-flex"
              >8</a
            >
            <a
              href="#"
              class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
              >9</a
            >
            <a
              href="#"
              class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
              >10</a
            >
            <a
              href="#"
              class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
            >
              <span class="sr-only">Next</span>
              <svg
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
                aria-hidden="true"
              >
                <path
                  fill-rule="evenodd"
                  d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z"
                  clip-rule="evenodd"
                />
              </svg>
            </a>
          </nav>
        </div>
      </div>
    </div>
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
          "PJ",
        ],
      ];

      for (const customer of customers) {
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
          customer.getPj(),
        ]);
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