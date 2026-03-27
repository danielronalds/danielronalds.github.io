import { ref } from 'vue';

export const useTerminal = () => {
  const path = ref('~');

  return {
    path,
  };
};
