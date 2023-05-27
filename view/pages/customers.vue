<template>
  <div>
    <form ref="file-input-form">
      <input type="file" @input="upload_excel($event)" />
    </form>
    <button
      class="bg-orange-500 text-sm active:bg-gray-700 cursor-pointer font-regular text-white px-4 py-2 rounded uppercase m-3"
      @click="export_excel"
    >
      Export
    </button>

    <div class="flex items-center">
      <div>
        <div class="relative mt-2 rounded-md shadow-sm">
          <input
            v-model="search"
            @input="
              page = 1;
              getCustomers();
            "
            type="text"
            name="price"
            class="block w-full rounded-md border-0 py-1.5 pl-2 pr-5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Cari"
          />
        </div>
      </div>
      <svg
        @click="toggleModalFilter"
        class="h-8 w-8 text-black ml-5"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3" />
      </svg>
    </div>
    <div
      v-if="modalFilter"
      class="modal fixed w-full h-full top-0 left-0 flex items-center justify-center"
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
            <div class="flex justify-around">
              <div class="pt-6">
                <span class="font-medium text-gray-900">Tahun/Bulan</span>
                <div>
                  <div class="relative mt-2 rounded-md shadow-sm">
                    <input
                      type="text"
                      name="price"
                      class="block w-full rounded-md border-0 py-1.5 pl-2 pr-5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                      placeholder="1444"
                    />
                    <div class="absolute inset-y-0 right-0 flex items-center">
                      <select
                        class="h-full rounded-md border-0 bg-transparent py-0 pl-2 pr-3 text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm"
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
              <div class="pt-6">
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
    <div class="w-full w-[200rem]">
      <div class="flex flex-row">
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[15rem] p-1"
        >
          Nama
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[5rem] p-1"
        >
          JK
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[5rem] p-1"
        >
          Status
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          KK
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Desa/Kampung
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Kol
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          PJ
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Bln Lalu
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Infaq
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Zakat
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Shadaaqah
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Ihsan
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Sabil
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Tabungan Fitrah
        </div>
        <div
          class="font-medium border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          Tabungan Qurban
        </div>
      </div>
      <div
        class="flex flex-row"
        v-for="customer in dataList"
        v-bind:key="customer._id"
      >
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[15rem] p-1"
        >
          {{ customer.nama }}
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[5rem] p-1"
        >
          {{ customer.jk }}
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[5rem] p-1"
        >
          {{ customer.status }}
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          {{ customer.nokk }}
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          {{ customer.kampung ? customer.kampung : customer.desakelurahan }}
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          {{ customer.kol }}
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          {{ customer.pj }}
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          TUNAI
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          <input
            type="number"
            name="price"
            class="block rounded-md w-32 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Infaq"
          />
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          <input
            type="number"
            name="price"
            class="block rounded-md w-32 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Zakat"
          />
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          <input
            type="number"
            name="price"
            class="block rounded-md w-32 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Shadaqah"
          />
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          <input
            type="number"
            name="price"
            class="block rounded-md w-32 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Ihsan"
          />
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          <input
            type="number"
            name="price"
            class="block rounded-md w-32 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Sabil"
          />
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          <input
            type="number"
            name="price"
            class="block rounded-md w-32 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Tabungan Fitrah"
          />
        </div>
        <div
          class="border border-inherit hover:border-indigo-300 basis-auto w-[10rem] p-1"
        >
          <input
            type="number"
            name="price"
            class="block rounded-md w-32 border-0 py-1.5 pl-2 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
            placeholder="Tabungan Qurban"
          />
        </div>
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
} from "../../proto/customers_pb";
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
    };

    let state: customerState = {
      page: 1,
      perPage: 30,
      totalPage: 1,
      dataList: [],
      modalFilter: false,
      search: "",
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
        payload.setPage(this.page);
        payload.setPerPage(this.perPage);
        payload.setSearch(this.search);
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
  },
});
</script>