<script setup lang="ts">
import type { TSource } from "@/types/emails";

defineProps<{ selectedEmail?: TSource }>();
</script>
<template>
  <div class="viewer" v-if="selectedEmail">
    <h3>{{ selectedEmail.Subject }} - {{ selectedEmail.Date }}</h3>

    <div class="filename" v-if="selectedEmail['X-Filename']">
      <p>File: {{ selectedEmail["X-Filename"] }}</p>
    </div>

    <p>
      {{ selectedEmail.Body }}
    </p>
  </div>
</template>

<style scoped>
h3 {
  font-size: 1.5rem;
  margin-bottom: 20px;
  font-weight: bold;
}

.filename {
  color: red;
  margin-bottom: 20px;
  width: max-content;
  border: 1px solid black;
  padding: 10px;
}

p {
  font-size: 1.2rem;
}
</style>
