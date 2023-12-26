<template>
  <div>
    <div class="flex items-center">
      <div>
        <div class="relative m-1 rounded-md shadow-sm">
          <input
            v-model="search"
            @input="
              page = 1;
              getCustomers();
            "
            type="text"
            name="price"
            class="block w-full rounded-md w-28 border-0 py-1.5 pl-2 pr-5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Cari"
          />
        </div>
      </div>

      <svg
        class="h-8 w-8 text-black"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        @click="toggleModalFilter"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z"
        />
      </svg>
    </div>
    <div
      v-if="modalFilter"
      class="modal fixed w-full h-full top-0 left-0 flex items-center justify-center z-10"
    >
      <div class="modal-overlay absolute w-full h-full bg-white opacity-100">
        <div class="modal-container fixed w-full h-full z-50 overflow-y-auto">
          <!-- Add margin if you want to see grey behind the modal-->
          <div class="modal-content container mx-auto h-auto text-left p-4">
            <!--Title-->
            <div class="flex justify-between items-center pb-2">
              <p class="text-2xl font-bold">Filter</p>
            </div>

            <!--Body-->
            <div class="flex flex-wrap justify-between">
              <div class="pr-6 pt-6 w-64">
                <span class="font-medium text-gray-900">Tahun/Bulan</span>
                <div>
                  <div class="relative mt-2 rounded-md shadow-sm">
                    <input
                      type="text"
                      name="price"
                      class="block w-full rounded-md w-28 border-0 py-1.5 pl-2 pr-5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                      placeholder="1444"
                    />
                    <div class="absolute inset-y-0 right-0 flex items-center">
                      <select
                        class="h-full rounded-md w-28 border-0 bg-transparent py-0 pl-2 pr-3 text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm"
                      >
                        <option value="1">Muharam</option>
                        <option value="2">Safar</option>
                        <option value="3">Rabiul Awal</option>
                        <option value="4">Rabiul Akhir</option>
                        <option value="5">Jumadil Awal</option>
                        <option value="6">Jumadil Akhir</option>
                        <option value="7">Rajab</option>
                        <option value="8">Syaban</option>
                        <option value="9">Ramadhan</option>
                        <option value="10">Syawal</option>
                        <option value="11">Dzulqadah</option>
                        <option value="12">Dzulhijah</option>
                      </select>
                    </div>
                  </div>
                </div>
              </div>
              <div class="pr-6 pt-6 w-48">
                <div class="space-y-3">
                  <span class="font-medium text-gray-900">Layanan Dasar</span>
                  <div class="flex items-center">
                    <input
                      id="filter-mobile-category-0"
                      name="category[]"
                      value="new-arrivals"
                      type="checkbox"
                      class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label
                      for="filter-mobile-category-0"
                      class="ml-3 min-w-0 flex-1 text-gray-500"
                      >Tunai</label
                    >
                  </div>
                  <div class="flex items-center">
                    <input
                      id="filter-mobile-category-1"
                      name="category[]"
                      value="sale"
                      type="checkbox"
                      class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label
                      for="filter-mobile-category-1"
                      class="ml-3 min-w-0 flex-1 text-gray-500"
                      >Terbina</label
                    >
                  </div>
                  <div class="flex items-center">
                    <input
                      id="filter-mobile-category-1"
                      name="category[]"
                      value="sale"
                      type="checkbox"
                      class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label
                      for="filter-mobile-category-1"
                      class="ml-3 min-w-0 flex-1 text-gray-500"
                      >Tidak Tunai</label
                    >
                  </div>
                  <div class="flex items-center">
                    <input
                      id="filter-mobile-category-1"
                      name="category[]"
                      value="sale"
                      type="checkbox"
                      class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label
                      for="filter-mobile-category-1"
                      class="ml-3 min-w-0 flex-1 text-gray-500"
                      >Tidak Terbina</label
                    >
                  </div>
                </div>
              </div>
              <div class="pr-6 pt-6 w-48">
                <div class="space-y-3">
                  <span class="font-medium text-gray-900">Tampilkan Kolom</span>
                  <div class="flex items-center">
                    <input
                      id="filter-mobile-category-0"
                      name="category[]"
                      value="new-arrivals"
                      type="checkbox"
                      class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label
                      for="filter-mobile-category-0"
                      class="ml-3 min-w-0 flex-1 text-gray-500"
                      >Keuangan</label
                    >
                  </div>
                  <div class="flex items-center">
                    <input
                      id="filter-mobile-category-1"
                      name="category[]"
                      value="sale"
                      type="checkbox"
                      class="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                    />
                    <label
                      for="filter-mobile-category-1"
                      class="ml-3 min-w-0 flex-1 text-gray-500"
                      >Bina</label
                    >
                  </div>
                </div>
              </div>
              <div class="pr-6 pt-6 w-48">
                <span class="font-medium text-gray-900">Import/Export Excel</span>
                <div>
                <button
                  class="bg-orange-500 text-sm active:bg-gray-700 cursor-pointer font-regular text-white px-4 py-2 rounded uppercase m-3"
                  @click="export_excel"
                >
                  Export
                </button>
                
                <form ref="file-input-form">
                  Import : <br/><input type="file" @input="upload_excel($event)" />
                </form>
                </div>
              </div>
            </div>

            <!--Footer-->
            <div class="flex justify-end pt-2 mt-10">
              <div
                @click="toggleModalFilter"
                class="px-4 bg-transparent p-3 rounded-lg text-indigo-500 hover:bg-gray-100 hover:text-indigo-400 mr-2"
              >
                Action
              </div>
              <div
                @click="toggleModalFilter"
                class="px-4 bg-indigo-500 p-3 rounded-lg text-white hover:bg-indigo-400"
              >
                Close
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="w-screen">
      <div id="demoDW">
        <table class="text-left text-sm font-light" id="demoDT">
          <thead class="border-b font-medium dark:border-neutral-500">
            <tr>
              <th scope="col" class="px-3 py-1 bg-cyan-100">Nama</th>
              <th scope="col" class="px-3 py-1 bg-cyan-100">JK</th>
              <th scope="col" class="px-3 py-1 bg-cyan-100">Status</th>
              <th scope="col" class="px-3 py-1 bg-cyan-100" @click="toggleSort('noKK')">
                <span>KK</span>
                <span v-if="sort.noKK === 1 && sort.noKK !== 0">⇑</span>
                <span v-else-if="sort.noKK !== 0">⇓</span>
              </th>
              <th scope="col" class="px-3 py-1 bg-cyan-100">Desa/Kampung</th>
              <th scope="col" class="px-3 py-1 bg-cyan-100">RT</th>
              <th scope="col" class="px-3 py-1 bg-cyan-100">RW</th>
              <th scope="col" class="px-3 py-1 bg-cyan-100">Kol</th>
              <th scope="col" class="px-3 py-1 bg-cyan-100" @click="toggleSort('pj')">
                <span>PJ</span>
                <span v-if="sort.pj === 1 && sort.pj !== 0">⇑</span>
                <span v-else-if="sort.pj !== 0">⇓</span>
              </th>
              <th scope="col" class="px-3 py-1 bg-red-100">BLN LALU</th>
              <th scope="col" class="px-3 py-1 bg-red-100">Infaq</th>
              <th scope="col" class="px-3 py-1 bg-red-100">Zakat</th>
              <th scope="col" class="px-3 py-1 bg-red-100">Shadaaqah</th>
              <th scope="col" class="px-3 py-1 bg-red-100">Ikhsan</th>
              <th scope="col" class="px-3 py-1 bg-red-100">Sabil</th>
              <th scope="col" class="px-3 py-1 bg-red-100">Tabungan Fitrah</th>
              <th scope="col" class="px-3 py-1 bg-red-100">Tabungan Qurban</th>
              <th scope="col" class="px-3 py-1 bg-lime-100">BLN LALU</th>
              <th scope="col" class="px-3 py-1 bg-lime-100">Terbina</th>
            </tr>
          </thead>
          <tbody>
            <tr
              class="border-b dark:border-neutral-500"
              v-for="customer in dataList"
              v-bind:key="customer._id"
            >
              <th
                scope="row"
                class="whitespace-nowrap max-w-[30%] px-3 py-2 font-medium bg-white"
              >
                {{ customer.nama }}
              </th>
              <td class="whitespace-nowrap px-3 py-2 font-medium">
                {{ customer.jk }}
              </td>
              <td class="whitespace-nowrap px-3 py-2 font-medium">
                {{ customer.status }}
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                {{ customer.nokk }}
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                {{
                  customer.kampung ? customer.kampung : customer.desakelurahan
                }}
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                {{ customer.rt }}
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                {{ customer.rw }}
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                {{ customer.kol }}
              </td>
              <td class="whitespace-nowrap px-3 py-2">{{ customer.pj }}</td>
              <td class="whitespace-nowrap px-3 py-2">TUNAI</td>
              <td class="whitespace-nowrap px-3 py-2">
                <input
                  @input="inputReguler($event, customer, 'infaq')"
                  step="1000"
                  type="number"
                  name="price"
                  class="block rounded-md w-28 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Infaq"
                />
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                <input
                  @input="inputReguler($event, customer, 'zakat')"
                  type="number"
                  step="1000"
                  name="price"
                  class="block rounded-md w-28 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Zakat"
                />
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                <input
                  @input="inputReguler($event, customer, 'shadaqah')"
                  type="number"
                  step="1000"
                  name="price"
                  class="block rounded-md w-28 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Shadaqah"
                />
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                <input
                  @input="inputReguler($event, customer, 'ikhsan')"
                  type="number"
                  step="1000"
                  name="price"
                  class="block rounded-md w-28 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Ikhsan"
                />
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                <input
                  @input="inputReguler($event, customer, 'sabil')"
                  type="number"
                  step="1000"
                  name="price"
                  class="block rounded-md w-28 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Sabil"
                />
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                <input
                  @input="inputReguler($event, customer, 'fitrah')"
                  type="number"
                  step="1000"
                  name="price"
                  class="block rounded-md w-28 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Fitrah"
                />
              </td>
              <td class="whitespace-nowrap px-3 py-2">
                <input
                  @input="inputReguler($event, customer, 'qurban')"
                  type="number"
                  step="1000"
                  name="price"
                  class="block rounded-md w-28 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="Qurban"
                />
              </td>
              <td class="whitespace-nowrap px-3 py-2">TERBINA</td>
              <td class="whitespace-nowrap px-3 py-2">
                <input type="checkbox" @input="inputReguler($event, customer, 'terbina')"  />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div
      class="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6"
    >
      <div class="flex flex-1 justify-between sm:hidden">
        <button
          @click="setPage(page > 1 ? page - 1 : 1)"
          class="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
        >
          Previous
        </button>
        <span class="font-medium">Page {{ page }} of {{ totalPage }}</span>
        <button
          @click="setPage(page < totalPage ? page + 1 : totalPage)"
          class="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
        >
          Next
        </button>
      </div>
      <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            Page
            <span class="font-medium">{{ page }}</span>
            of
            <span class="font-medium">{{ totalPage }}</span>
          </p>
        </div>
        <div>
          <nav
            class="isolate inline-flex -space-x-px rounded-md shadow-sm"
            aria-label="Pagination"
          >
            <button
              @click="setPage(page > 1 ? page - 1 : 1)"
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
            </button>
            <!-- Current: "z-10 bg-indigo-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600", Default: "text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0" -->
            <button
              @click="setPage(currentPage)"
              v-for="currentPage of totalPage"
              v-if="currentPage < page + 3 && currentPage > page - 3"
              v-bind:key="currentPage"
              :class="`${
                currentPage === page
                  ? 'bg-indigo-600 text-white'
                  : 'text-gray-900'
              } relative inline-flex items-center px-4 py-2 text-sm font-semibold ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0`"
            >
              {{ currentPage }}
            </button>
            <button
              @click="setPage(page < totalPage ? page + 1 : totalPage)"
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
            </button>
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
  CustomerSortParam
} from "../../proto/customers_pb";
import {
  RegulerCreateRequest
} from "../../proto/reguler_pb";
import * as XLSX from "xlsx";

let getTimeout: any = null;

export default Vue.extend({
  name: "Customer",
  data: () => {
    type customerState = {
      page: number;
      perPage: number;
      totalPage: number;
      dataList: Array<Customer.AsObject>;
      modalFilter: boolean;
      search: string;
      sort: {[key:string]: number}
    };

    let state: customerState = {
      page: 1,
      perPage: 30,
      totalPage: 1,
      dataList: [],
      modalFilter: false,
      search: "",
      sort: {
        noKK: 0,
        pj: 0
      }
    };

    return state;
  },
  async mounted() {
    this.getCustomers();
  },
  methods: {
    async getCustomers() {
      if (getTimeout) clearTimeout(getTimeout);
      getTimeout = setTimeout(async () => {
        let payload = new CustomerGetRequest();
        let sortPaylad = new CustomerSortParam()
        payload.setPage(this.page);
        payload.setPerPage(this.perPage);
        payload.setSearch(this.search);
        if(this.sort.noKK){
          sortPaylad.setNokk(this.sort.noKK)
        }
        if(this.sort.pj){
          sortPaylad.setPj(this.sort.pj)
        }
        payload.setSort(sortPaylad)
        let response = await (<Promise<CustomerGetResponse>>(
          this.$store.dispatch("getCustomers", payload)
        ));

        if (response) {
          this.dataList = [];
          let obj = response.toObject();
          this.totalPage = obj.totalPage;
          obj.dataList.map((customer: Customer.AsObject) => {
            this.dataList.push(customer);
          });
        }
      }, 200);
    },
    toggleSort(column: string){
      if(!this.sort?.[column]){
        this.sort[column] = 1
      }else if(this.sort?.[column] === 1){
        this.sort[column] = -1
      }else {
        this.sort[column] = 0
      }

      this.getCustomers()
    },
    toggleModalFilter() {
      this.modalFilter = !this.modalFilter;
    },
    setPage(page: number) {
      this.page = page;
      this.getCustomers();
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
          customer.setJk(customerData["JK"]);
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
          "JK",
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
          customer.getJk(),
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

    async inputReguler(val: any, customerData: Customer.AsObject, columnName: string){
      console.log('value', val.target.checked, customerData)
      let regulerCreateRequest = new RegulerCreateRequest() 
      regulerCreateRequest.setBulan("1")
      regulerCreateRequest.setTahun("1444")
      regulerCreateRequest.setCustomerid(customerData.id)

      if(columnName === 'infaq'){
        regulerCreateRequest.setInfaq(val.target.value)
      }

      if(columnName === 'zakat'){
        regulerCreateRequest.setZakat(val.target.value)
      }

      if(columnName === 'shadaqah'){
        regulerCreateRequest.setShadaqah(val.target.value)
      }

      if(columnName === 'ikhsan'){
        regulerCreateRequest.setIkhsan(val.target.value)
      }

      if(columnName === 'sabil'){
        regulerCreateRequest.setSabil(val.target.value)
      }

      if(columnName === 'qurban'){
        regulerCreateRequest.setTabunganqurban(val.target.value)
      }

      if(columnName === 'fitrah'){
        regulerCreateRequest.setTabunganfitrah(val.target.value)
      }

      if(columnName === 'terbina'){
        regulerCreateRequest.setBina(val.target.checked ? "TERBINA" : "PASIF")
      }

      this.$store.dispatch("addReguler", regulerCreateRequest);
    }
  },
});
</script>