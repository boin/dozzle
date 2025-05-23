<template>
  <ScrollableView :scrollable="scrollable" v-if="service.name">
    <template #header>
      <div class="mx-2 flex items-center gap-2 md:ml-4">
        <ph:stack-simple />
        <ContainerDropdown :containers="service.containers">{{ service.name }}</ContainerDropdown>
        <MultiContainerStat class="ml-auto" :containers="service.containers" />
        <MultiContainerActionToolbar class="max-md:hidden" @clear="viewer?.clear()" />
      </div>
    </template>
    <template #default>
      <ViewerWithSource
        ref="viewer"
        :stream-source="useServiceStream"
        :entity="service"
        :visible-keys="new Map<string[], boolean>()"
      />
    </template>
  </ScrollableView>
</template>

<script lang="ts" setup>
import { Service } from "@/models/Stack";
import ViewerWithSource from "@/components/LogViewer/ViewerWithSource.vue";
import { ComponentExposed } from "vue-component-type-helpers";

const { name, scrollable = false } = defineProps<{
  scrollable?: boolean;
  name: string;
}>();

const viewer = ref<ComponentExposed<typeof ViewerWithSource>>();
const store = useSwarmStore();
const { services } = storeToRefs(store) as unknown as { services: Ref<Service[]> };
const service = computed(() => services.value.find((s) => s.name === name) ?? new Service("", []));

provideLoggingContext(
  toRef(() => service.value.containers),
  { showContainerName: true, showHostname: false },
);
</script>
