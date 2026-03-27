<script setup lang="ts">
import { SquareTerminal } from 'lucide-vue-next';

defineEmits<{ click: [] }>();
</script>
<template>
  <button
    class="rounded-lg bg-[#11111b] text-cp-bg0 shadow-md size-14 cursor-pointer flex justify-center items-center"
    @click="$emit('click')"
  >
    <SquareTerminal :size="36" />
  </button>
</template>
