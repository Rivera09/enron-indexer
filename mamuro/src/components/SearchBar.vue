<script setup lang="ts">
import { ref } from "vue";

const emit = defineEmits(["onSearch", "reset"]);
const term = ref("");

const submitForm = () => {
  emit("onSearch", term.value);
};

const resetTable = () => {
  emit("reset");
};
</script>

<template>
  <form @submit.prevent="submitForm" class="mb-5 flex gap-x-5">
    <input
      type="text"
      v-model="term"
      name="term"
      placeholder="Search email"
      autocomplete="off"
      class="font-semibold rounded border-2 border-indigo-900 flex-1 p-2.5 outline-none"
    />
    <input
      type="button"
      @click="resetTable"
      value="reset"
      class="p-2.5 rounded border-2 border-indigo-900 duration-300 hover:bg-indigo-900 hover:text-white"
    />
    <button
      type="submit"
      class="p-2.5 rounded border-2 border-indigo-900 duration-300 hover:bg-indigo-900 hover:text-white"
    >
      Search
    </button>
  </form>
</template>
