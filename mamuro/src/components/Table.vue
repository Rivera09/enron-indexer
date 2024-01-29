<script setup lang="ts">
import type { TSource } from "@/types/emails";

defineProps<{ emailsList: TSource[]; selectedId?: string }>();
const emit = defineEmits(["clicked"]);

const handleClick = (id: string) => {
  emit("clicked", id);
};
</script>

<template>
  <table class="border-collapse self-start">
    <thead>
      <tr>
        <th class="p-1 border-2 border-black">Subject</th>
        <th class="p-1 border-2 border-black">From</th>
        <th class="p-1 border-2 border-black">To</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="{ Subject, From, To, id } in emailsList"
        :class="{ selected: !!selectedId && selectedId === id }"
        class="hover:bg-indigo-200 duration-300 cursor-pointer"
        @click="handleClick(id)"
      >
        <th class="p-1 text-left border-2 border-black">
          {{ Subject || "No subject" }}
        </th>
        <th class="p-1 text-left border-2 border-black">{{ From }}</th>
        <th class="p-1 text-left border-2 border-black">{{ To || "" }}</th>
      </tr>
    </tbody>
  </table>
</template>

<style scoped>
tr.selected {
  background-color: #444888;
  color: white;
}
</style>
