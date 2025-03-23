<template>
  <div ref="optionRef" class="x-select-option" :class="{ 'x-select-option-selected': value == selectedValue }"
    @click="handleClick">
    <slot>{{ label }}</slot>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, inject } from 'vue'
const props = defineProps({
  value: { type: String, required: false },
  label: { type: String, required: false }
})
const optionRef = ref(null)
const textContent = ref('');
const value = computed(() => props.value ?? props.label ?? textContent.value)
const label = computed(() => props.label ?? textContent.value)
onMounted(() => {
  if (!props.label && optionRef.value) {
    const text = optionRef.value.textContent ?? '';
    if (textContent.value !== text) {
      textContent.value = text;
    }
  }
  registerOption(value.value, label.value)
})
const selectedValue = inject('selectSelectedValue')
const updateSelectedValue = inject('updateSelecteSelectedValue')
const registerOption = inject('registerSelectOption')
const handleClick = () => {
  updateSelectedValue(value.value, label.value)
}
</script>

<style lang="less">
.x-select-option {
  height: 1.5rem;
  display: flex;
  align-items: center;
  border-radius: .25rem;
  padding: 0 .5rem;

  &:hover {
    background-color: var(--color-primary-2);
  }
}
.x-select-option-selected{
  background-color: var(--color-primary-2);
}
</style>