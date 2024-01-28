<script setup lang="ts">
import axios from "axios";
import { ref } from "vue";
import type { Ref } from "vue";
import type { TSource } from "@/types/emails";
import SearchBar from "@/components/SearchBar.vue";
import Header from "@/components/Header.vue";
import Table from "@/components/Table.vue";

const emails: Ref<TSource[]> = ref([]);

const handleCreate = async (term: string) => {
  const res = await axios.get<TSource[]>("http://localhost:3000/email");
  emails.value = res.data;
};
</script>

<template>
  <main>
    <Header />
    <SearchBar @created="handleCreate" />
    <div class="flex-c">
      <Table :emails-list="emails" />
      <div class="viewer">
        <h3>Subject</h3>
        <p>
          Lorem ipsum, dolor sit amet consectetur adipisicing elit. Distinctio
          impedit corporis delectus animi eum optio eos ea expedita, odio dicta
          necessitatibus dolorem vero maiores repellendus commodi. Corporis
          cumque tenetur quasi.
        </p>
        <div v-for="email in emails">
          <p>{{ email._source.Subject }}</p>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.flex-c {
  display: flex;
  gap: 20px;
}

.flex-c > * {
  flex: 1;
}
</style>
