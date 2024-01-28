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
};

const handleClick = (id: string) => {
  const emailData = emails.value.find((e) => e.id === id);
  if (emailData) selectedEmail.value = emailData;
};
</script>

<template>
  <Header />
  <main>
    <SearchBar @on-search="handleSearch" @reset="handleReset" />
    <div class="flex-c">
      <Table
        :emails-list="emails"
        :selected-id="selectedEmail?.id"
        @clicked="handleClick"
      />
      <EmailReader :selected-email="selectedEmail" />
    </div>
  </main>
</template>

<style scoped>
main {
  padding: 20px;
}

.flex-c {
  display: flex;
  gap: 20px;
}

.flex-c > * {
  flex: 1;
}
</style>
