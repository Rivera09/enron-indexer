<script setup lang="ts">
import type { TSource } from "@/types/emails";

defineProps<{ selectedEmail?: TSource }>();
</script>
<template>
  <div class="viewer" v-if="selectedEmail">
    <h3 class="text-2xl mb-5 font-bold">
      {{ selectedEmail.Subject }} - {{ selectedEmail.Date }} (from:
      {{ selectedEmail.From }})
    </h3>

    <p
      class="text-indigo-900 rounded border-2 border-indigo-900 font-medium p-2.5 max-w-max mb-5 text-lg"
      v-if="selectedEmail['X-Filename']"
    >
      File: {{ selectedEmail["X-Filename"] }}
    </p>

    <p class="text-lg">
      {{ selectedEmail.Body }}
    </p>
  </div>
</template>
