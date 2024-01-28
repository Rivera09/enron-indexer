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
    <input type="button" @click="resetTable" value="reset" />
    <button type="submit">Search</button>
  </form>
</template>

<style scoped>
form {
  margin: 20px 0;
  display: flex;
  gap: 20px;
}

input {
  font-weight: 600;
}

button,
input[type="button"] {
  padding: 10px;
  border-radius: 4px;
  border: 1px solid black;
  transition: 0.3s;
  width: max-content;
  flex: 0;
  cursor: pointer;
}

button:hover,
input[type="button"]:hover {
  background-color: #444888;
  color: white;
}

input {
  border: 1px solid black;
  flex: 1;
  max-width: 100%;
  padding: 10px;
  outline: none;
  border-radius: 4px;
}
</style>
