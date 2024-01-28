<script setup lang="ts">
import type { TSource } from "@/types/emails";

defineProps<{ emailsList: TSource[]; selectedId?: string }>();
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
        :class="{ selected: !!selectedId && selectedId === id }"
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
  border-collapse: collapse;
  align-self: flex-start;
}

td,
th {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr {
  cursor: pointer;
  transition: 0.3s;
}

tr:hover {
  background-color: #f0edff;
}

tr.selected {
  background-color: #444888;
  color: white;
}
</style>
