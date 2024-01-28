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
  <form @submit.prevent="submitForm">
    <input
      type="text"
      v-model="term"
      name="term"
      placeholder="Search email"
      autocomplete="off"
    />
    <button @click="resetTable">Reset</button>
    <button type="submit">Search</button>
  </form>
</template>

<style scoped>
form {
  margin: 20px 0;
  display: flex;
  gap: 20px;
}

button {
  padding: 10px;
  border-radius: 8px;
  border: 1px solid black;
}

input {
  border: 1px solid black;
  flex: 1;
  max-width: 100%;
  padding: 10px;
  outline: none;
  border-radius: 8px;
}
</style>
