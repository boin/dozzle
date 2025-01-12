<template>
  <LogItem :logEntry>
    <div class="whitespace-pre-wrap" :data-event="logEntry.event" v-html="logEntry.message"></div>
  </LogItem>

  <div
    class="alert alert-info my-4 w-auto flex-none font-sans text-[1rem] md:mx-auto md:w-1/2"
    v-if="followEligible && showCard"
  >
    <carbon:information class="size-6 shrink-0 stroke-current" />
    <div>
      <h3 class="text-lg font-bold">{{ $t("alert.similar-container-found.title") }}</h3>
      {{ $t("alert.similar-container-found.message", { containerId: nextContainer.id }) }}
    </div>
    <div>
      <TimedButton
        v-if="automaticRedirect"
        class="btn-primary btn-sm"
        @finished="redirectNow()"
        @click="showCard = false"
      >
        {{ $t("button.cancel") }}
      </TimedButton>
      <router-link
        :to="{ name: '/container/[id]', params: { id: nextContainer.id } }"
        class="btn btn-primary btn-sm"
        v-else
      >
        {{ $t("button.redirect") }}
      </router-link>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ContainerEventLogEntry } from "@/models/LogEntry";
const router = useRouter();
const { showToast } = useToast();
const { t } = useI18n();

const { logEntry } = defineProps<{
  logEntry: ContainerEventLogEntry;
  showContainerName?: boolean;
}>();

const showCard = ref(true);
const { containers } = useLoggingContext();
const store = useContainerStore();
const { containers: allContainers } = storeToRefs(store);

const nextContainer = computed(
  () =>
    [
      ...allContainers.value.filter(
        (c) => c.created > containers.value[0].created && c.name === containers.value[0].name,
      ),
    ].sort((a, b) => +a.created - +b.created)[0],
);

const followEligible = computed(
  () =>
    router.currentRoute.value.name === "/container/[id]" && // we are on a container page
    logEntry.event === "container-stopped" && // container was stopped
    containers.value.length === 1 && // only one container
    Date.now() - +logEntry.date < 5 * 60 * 1000 && // was stopped in the last 5 minutes
    nextContainer.value !== undefined, // there is a next container
);

function redirectNow() {
  showToast(
    {
      title: t("alert.redirected.title"),
      message: t("alert.redirected.message", { containerId: nextContainer.value?.id }),
      type: "info",
    },
    { expire: 5000 },
  );
  router.push({ name: "/container/[id]", params: { id: nextContainer.value?.id } });
}
</script>

<style scoped>
@import "@/main.css" reference;
[data-event="container-stopped"] {
  @apply text-red;
}
[data-event="container-started"] {
  @apply text-green;
}
</style>
