<template>
  <div class="form-item">
    <div class="form-item-label">{{ props.label }}</div>
    <div class="form-item-input">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  label: { type: String, required: true }
})

</script>

<style lang="less">
.form-item {
  height: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
</style>
