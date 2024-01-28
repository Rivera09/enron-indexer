<script setup lang="ts">
import type { TSource } from "@/types/emails";

defineProps<{ selectedEmail?: TSource }>();
</script>
<template>
  <div class="viewer" v-if="selectedEmail">
    <h3>
      {{ selectedEmail.Subject }} - {{ selectedEmail.Date }} (from:
      {{ selectedEmail.From }})
    </h3>

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
  color: #444888;
  font-weight: 500;
  margin-bottom: 20px;
  width: max-content;
  border: 1px solid #444888;
  padding: 10px;
  border-radius: 4px;
}

p {
  font-size: 1.2rem;
}
</style>
