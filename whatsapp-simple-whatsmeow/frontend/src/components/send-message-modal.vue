<template>
  <dialog id="send-message-modal" class="modal">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Kirim Pesan</h3>

      <div class="mt-5 w-full flex justify-center items-center">
        <div class="w-full flex flex-col gap-2">
          <label class="form-control w-full">
            <div class="label">
              <span class="label-text">Nomor Penerima</span>
            </div>
            <input
              type="text"
              placeholder="6281234567890"
              class="input input-bordered w-full"
              v-model="state.receiver"
            />
          </label>

          <label class="form-control">
            <div class="label">
              <span class="label-text">Pesan Teks</span>
            </div>
            <textarea
              class="textarea textarea-bordered h-24"
              placeholder="Pesan Teks"
              v-model="state.text"
            ></textarea>
          </label>

          <div class="w-full mt-3">
            <button
              class="btn btn-success btn-block btn-sm"
              :disabled="!state.receiver || !state.text || state.loading"
              @click="handleClickSend"
            >
              <span class="loading" v-if="state.loading"></span>
              Kirim
            </button>
          </div>
        </div>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button @click.prevent="handleCloseModal">close</button>
    </form>
  </dialog>
</template>

<script setup>
import { defineProps, defineEmits, reactive } from "vue";
import { sendMessage } from "@/services/message.service";

const state = reactive({
  receiver: "",
  text: "",
  loading: false,
});

const emits = defineEmits(["close"]);
const props = defineProps({
  number: {
    type: Object,
    default: () => ({}),
  },
});

async function handleClickSend() {
  if (state.loading) return;
  if (!state.receiver) return;
  if (!state.text) return;

  state.loading = true;

  try {
    const { data } = await sendMessage(
      props.number._id,
      state.receiver,
      state.text
    );

    alert(`pesan berhasil dikirim dengan id ${data.messageId}`);
  } catch (error) {
    alert(error?.response?.data?.errorMessage || error?.message);
  }

  state.loading = false;
  state.receiver = "";
  state.text = "";
}

function handleCloseModal() {
  emits("close");
}
</script>
