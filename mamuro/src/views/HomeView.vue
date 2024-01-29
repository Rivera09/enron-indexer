<script setup lang="ts">
import axios from "axios";
import { ref } from "vue";
import type { Ref } from "vue";
import type { TEmailResponse, TSource } from "@/types/emails";
import SearchBar from "@/components/SearchBar.vue";
import Header from "@/components/Header.vue";
import Table from "@/components/Table.vue";
import EmailReader from "@/components/EmailReader.vue";

const emails: Ref<TSource[]> = ref([]);
const selectedEmail = ref<TSource>();

const handleSearch = async (term: string) => {
  if (!term) return;
  const res = await axios.get<TEmailResponse[]>(
    `http://localhost:3000/email?term=${term}`
  );
  if (res.status === 200) emails.value = res.data.map((d) => d._source);
};

const handleReset = () => {
  emails.value = [];
  selectedEmail.value = undefined;
};

const handleClick = (id: string) => {
  const emailData = emails.value.find((e) => e.id === id);
  if (emailData) selectedEmail.value = emailData;
};
</script>

<template>
  <Header />
  <main class="p-5">
    <SearchBar @on-search="handleSearch" @reset="handleReset" />
    <div class="flex gap-5">
      <Table
        :emails-list="emails"
        :selected-id="selectedEmail?.id"
        @clicked="handleClick"
        v-if="emails.length"
        class="flex-1"
      />
      <p v-else class="text-center flex-1 mt-10 text-3xl font-bold">
        Email list is empty
      </p>
      <EmailReader :selected-email="selectedEmail" class="flex-1" />
    </div>
  </main>
</template>
