<script setup lang="ts">
import { ref } from "vue";

const emit = defineEmits(["created"]);
const term = ref("");

const submitForm = () => {
  emit("created", term.value);
  // Your form submission logic goes here
  // console.log("Form submitted!", event.target);
  // console.log("Password:", this.searchInput);

  // You can call other functions or perform actions as needed
};
</script>

<template>
  <form @submit.prevent="submitForm">
    <input
      type="text"
      v-model="term"
      name="term"
      placeholder="Search email"
      autocomplete="false"
    />
    <button type="submit">Search</button>
  </form>
</template>

<!-- <script>
export default {
  data() {
    return {};
  },
  methods: {
    submitForm() {
      const emit = defineEmits(["created"]);
      emit("created");
      // Your form submission logic goes here
      console.log("Form submitted!");
      console.log("Password:", this.searchInput);

      // You can call other functions or perform actions as needed
    },
  },
};
</script> -->

<style scoped>
form {
  margin: 20px 10px;
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
  /* margin: 20px 10px; */
  /* width: 100%; */
  flex: 1;
  max-width: 100%;
  padding: 10px;
  outline: none;
  border-radius: 8px;
}
</style>
