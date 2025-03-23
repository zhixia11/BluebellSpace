<template>
  <div class="form-group">
    <div class="form-group-header">
      <div class="form-group-header-left">
        <span class="form-group-header-icon">
          <slot name="icon"></slot>
        </span>
        <span class="form-group-header-label" :title="props.help">{{ props.label }}</span>
      </div>
      <div class="form-group-header-right">
        <div class="form-group-header-ext" v-if="enable != false">
          <slot name="ext"></slot>
        </div>
        <div class="form-group-header-switch" v-if="enable != undefined && enable != false">
          <a-switch type="line" size="small" v-model="enable" />
        </div>
      </div>
    </div>
    <div class="form-group-content" v-if="enable != false">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  label: { type: String, requered: true },
  help: { type: String, requered: false }
})
const enable = defineModel({ type: Boolean, default: undefined })
</script>

<style lang="less">
.form-group {
  width: 100%;
  margin-bottom: .5rem;
}

.form-group-header {
  height: 1.75rem;
  font-size: .875rem;
  display: flex;
}

.form-group-header-left {
  display: flex;
  align-items: center;
  line-height: 100%;
}

.form-group-header-right {
  flex: 1;
  display: flex;
  justify-content: end;
}

.form-group-header-label {
  display: flex;
  align-items: center;
  font-weight: bold;
  border-bottom: 1px solid var(--color-primary-1);
  padding: 1px;
}

.form-group-header-ext {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: right;

  &>.x-button {
    padding: .375rem .5rem !important;
  }
}

.form-group-header-switch {
  width: fit-content;
  display: flex;
  align-items: center;
  margin-left: .25rem;
  margin-right: 0;
}

.form-group-content {
  .form-item:not(:last-child) {
    margin-bottom: .375rem;
  }
}
</style>
