<template>
  <div class="data-list" :class="{ 'data-list-wrap': wrap }">
    <div class="data-list-item" v-for="(val, index) in data" :key="val" :title="val" v-if="Array.isArray(data)">
      <div class="data-list-item-content">{{ val }}</div>
      <div class="data-list-item-close" @click="remove(index)">
        <icon name="close" size="nano" />
      </div>
    </div>
    <div class="data-list-item" v-if="!Array.isArray(data) && !!data">
      <div class="data-list-item-content">{{ data }}</div>
      <div class="data-list-item-close" @click="remove">
        <icon name="close" size="nano" />
      </div>
    </div>
  </div>
</template>

<script setup>
import Icon from '../icon/Icon.vue'
import Modal from '../dialog'
const props = defineProps({
  wrap: {
    type: Boolean,
    default: false
  }
})
const data = defineModel()

const remove = (index) => {
  Modal.show({
    message: '确定要删除吗',
    width: 280,
    height: 140,
    onOk: () => {
      if (Array.isArray(data.value)) {
        data.value.splice(index, 1)
      } else {
        data.value = ''
      }
    }
  })
}

</script>

<style lang="less">
.data-list {
  user-select: none;
}

.data-list-item {
  height: 1rem;
  padding: .25rem;
  color: #6B7785;
  border-radius: .25rem;
  cursor: default;
  font-size: .75rem;
  display: flex;
  flex-flow: row;

  .data-list-item-content {
    flex: 1;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    display: flex;
    align-items: center;
  }

  .data-list-item-close {
    padding-left: .25rem;
    border-radius: .25rem;
    display: none;
    justify-content: center;
    align-items: center;
    cursor: pointer;

    &:hover {
      color: #C396ED;
    }

    &:active {
      color: #ffffff;
    }
  }

  &:hover {
    background-color: #e6ddfa !important;

    .data-list-item-close {
      display: flex;
      margin: 0 .125rem;
    }
  }
}

.data-list-wrap {
  display: flex;
  flex-wrap: wrap;
  gap: .375rem;

  .data-list-item {
    padding: .25rem .5rem;
    background-color: #F2F3F5;
  }

  .data-list-item-close {
    display: flex;
    padding: 0;
    margin-left: .125rem;
  }
}
</style>
