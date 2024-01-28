<script setup lang="ts">
import type { TSource } from "@/types/emails";

defineProps<{ emailsList: TSource[] }>();
const emit = defineEmits(["clicked"]);

const handleClick = (id: string) => {
  emit("clicked", id);
};
</script>

<template>
  <table>
    <thead>
      <tr>
        <th>Subject</th>
        <th>From</th>
        <th>To</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="{ Subject, From, To, id } in emailsList"
        @click="handleClick(id)"
      >
        <th>{{ Subject || "No subject" }}</th>
        <th>{{ From }}</th>
        <th>{{ To || "" }}</th>
      </tr>
    </tbody>
  </table>
</template>

<style scoped>
table {
  /* margin: auto; */

  /* width: 95%; */
  border-collapse: collapse;
  align-self: flex-start;
}

td,
th {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}
</style>
