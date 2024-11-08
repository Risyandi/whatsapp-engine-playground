<template>
  <dialog id="scan-qr-modal" class="modal">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Scan Kode QR</h3>

      <div class="mt-5 w-full flex justify-center items-center">
        <div v-if="props.qrcode">
          <img :src="`data:image/png;base64,${props.qrcode}`" class="w-64" />
        </div>
        <div v-else>
          <div class="loading"></div>
        </div>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button @click.prevent="handleCloseModal">close</button>
    </form>
  </dialog>
</template>

<script setup>
import { defineProps, defineEmits } from "vue";

const emits = defineEmits(["close"]);
const props = defineProps({
  qrcode: {
    type: String,
    default: "",
  },
});

function handleCloseModal() {
  emits("close");
}
</script>
