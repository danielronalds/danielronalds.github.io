<script setup lang="ts">
import { Hand } from 'lucide-vue-next';

defineEmits<{ click: [] }>();
</script>
<template>
  <button
    class="p-2 rounded-lg bg-cp-orange shadow-md size-14 cursor-pointer flex justify-center items-center"
    @click="$emit('click')"
  >
    <Hand :size="30" />
  </button>
</template>
