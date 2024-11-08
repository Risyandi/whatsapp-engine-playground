<template>
  <div class="w-full h-screen mx-auto max-w-2xl bg-base-200 shadow-lg">
    <div class="w-full p-5 text-2xl">wa-message-example</div>

    <div class="px-5 py-2 w-full flex flex-col gap-3">
      <div class="w-full mb-2 flex justify-end">
        <button class="btn btn-success btn-sm" @click="handleClickAdd">
          Tambah WhatsApp
        </button>
      </div>

      <template v-if="state.numbers.length">
        <div
          v-for="(number, index) in state.numbers"
          :key="index"
          class="w-full border px-5 py-4 rounded-md flex gap-3 items-center justify-between"
          :class="{
            'border-red-400': !number.isConnected,
            'border-green-400': number.isConnected,
          }"
        >
          <div class="w-full flex gap-3 items-center">
            <div>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                fill="currentColor"
                class="bi bi-whatsapp"
                viewBox="0 0 16 16"
              >
                <path
                  d="M13.601 2.326A7.85 7.85 0 0 0 7.994 0C3.627 0 .068 3.558.064 7.926c0 1.399.366 2.76 1.057 3.965L0 16l4.204-1.102a7.9 7.9 0 0 0 3.79.965h.004c4.368 0 7.926-3.558 7.93-7.93A7.9 7.9 0 0 0 13.6 2.326zM7.994 14.521a6.6 6.6 0 0 1-3.356-.92l-.24-.144-2.494.654.666-2.433-.156-.251a6.56 6.56 0 0 1-1.007-3.505c0-3.626 2.957-6.584 6.591-6.584a6.56 6.56 0 0 1 4.66 1.931 6.56 6.56 0 0 1 1.928 4.66c-.004 3.639-2.961 6.592-6.592 6.592m3.615-4.934c-.197-.099-1.17-.578-1.353-.646-.182-.065-.315-.099-.445.099-.133.197-.513.646-.627.775-.114.133-.232.148-.43.05-.197-.1-.836-.308-1.592-.985-.59-.525-.985-1.175-1.103-1.372-.114-.198-.011-.304.088-.403.087-.088.197-.232.296-.346.1-.114.133-.198.198-.33.065-.134.034-.248-.015-.347-.05-.099-.445-1.076-.612-1.47-.16-.389-.323-.335-.445-.34-.114-.007-.247-.007-.38-.007a.73.73 0 0 0-.529.247c-.182.198-.691.677-.691 1.654s.71 1.916.81 2.049c.098.133 1.394 2.132 3.383 2.992.47.205.84.326 1.129.418.475.152.904.129 1.246.08.38-.058 1.171-.48 1.338-.943.164-.464.164-.86.114-.943-.049-.084-.182-.133-.38-.232"
                />
              </svg>
            </div>
            <div>
              <div class="text-sm">{{ number.name || "-" }}</div>
              <div class="text-xs text-zinc-500">
                {{ number.phoneNumber || "-" }}
              </div>
            </div>
          </div>
          <div class="dropdown dropdown-end">
            <div tabindex="0" role="button" class="">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-three-dots-vertical"
                viewBox="0 0 16 16"
              >
                <path
                  d="M9.5 13a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0m0-5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0m0-5a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0"
                />
              </svg>
            </div>
            <ul
              tabindex="0"
              class="dropdown-content menu bg-base-100 rounded-box z-[1] w-52 p-2 shadow"
            >
              <li v-if="!number.isConnected" @click="handleClickScanQR(number)">
                <a>Scan Kode QR</a>
              </li>
              <li
                v-if="number.isConnected"
                @click="handleClickSendMessage(number)"
              >
                <a>Kirim Pesan</a>
              </li>
              <li @click="handleClickDelete(number)">
                <a class="text-red-500">Hapus Nomor</a>
              </li>
            </ul>
          </div>
        </div>
      </template>
      <template v-else>
        <div class="w-full flex justify-center items-center mt-12">
          <div>Belum ada Nomor</div>
        </div>
      </template>
    </div>
  </div>

  <scan-qr-modal :qrcode="state.selected?.qrcode" @close="handleCloseModal" />
  <send-message-modal :number="state.selected" @close="handleCloseModal" />
</template>

<script setup>
import { onMounted, reactive } from "vue";
import {
  getNumber,
  getOneNumber,
  addNumber,
  removeNumber,
} from "@/services/number.service";

import scanQrModal from "@/components/scan-qr-modal.vue";
import sendMessageModal from "@/components/send-message-modal.vue";

const state = reactive({
  numbers: [],
  selected: null,
});

async function handleClickScanQR(number) {
  state.selected = number;

  document.getElementById("scan-qr-modal").showModal();

  while (state.selected) {
    try {
      const { data } = await getOneNumber(number._id);
      state.selected = data;

      if (data.isConnected) {
        handleGetNumbers();
        handleCloseModal();
      }
    } catch (error) {
      //
    }

    await delay(3000);
  }
}

async function handleClickAdd() {
  try {
    const { data } = await addNumber();
    handleClickScanQR(data);
    handleGetNumbers();
  } catch (error) {
    alert(error?.response?.data?.errorMessage || error?.message);
  }
}

async function handleClickDelete(number) {
  try {
    await removeNumber(number._id);
    handleGetNumbers();
  } catch (error) {
    alert(error?.response?.data?.errorMessage || error?.message);
  }
}

async function handleGetNumbers() {
  try {
    const { data } = await getNumber();
    state.numbers = data.numbers;
  } catch (error) {
    alert(error?.response?.data?.errorMessage || error?.message);
  }
}

function handleClickSendMessage(number) {
  state.selected = number;
  document.getElementById("send-message-modal").showModal();
}

function handleCloseModal() {
  state.selected = null;
  document.getElementById("scan-qr-modal").close();
  document.getElementById("send-message-modal").close();
}

function delay(milisecond) {
  return new Promise((resolve) => setTimeout(resolve, milisecond));
}

onMounted(() => {
  handleGetNumbers();
});
</script>
