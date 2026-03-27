<script setup lang="ts">
import type { Ref } from 'vue';
import { ChevronRight } from 'lucide-vue-next';

defineProps<{ path: Ref<string> }>();
</script>

<template>
  <div class="flex flex-row gap-0 text-cp-aqua items-center">
    <span class="font-bold">{{ path }}</span>
    <ChevronRight :size="20" />
  </div>
</template>
