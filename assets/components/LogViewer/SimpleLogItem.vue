<template>
  <LogItem :logEntry>
    <div
      class="[word-break:break-word] whitespace-pre-wrap group-[.disable-wrap]:whitespace-nowrap"
      v-html="colorize(logEntry.message)"
    ></div>
    <LogMessageActions
      class="absolute -right-1 opacity-0 transition-opacity delay-150 duration-250 group-hover/entry:opacity-100"
      :message="() => stripAnsi(logEntry.rawMessage)"
      :log-entry="logEntry"
      v-if="containers.length > 0"
    />
  </LogItem>
</template>
<script lang="ts" setup>
import { SimpleLogEntry } from "@/models/LogEntry";
import AnsiConvertor from "ansi-to-html";
import stripAnsi from "strip-ansi";
const { containers } = useLoggingContext();

const ansiConvertor = new AnsiConvertor({
  escapeXML: false,
  fg: "var(--color-base-content)",
  bg: "var(--color-base-100)",
});

defineProps<{
  logEntry: SimpleLogEntry;
}>();

const colorize = (value: string) => ansiConvertor.toHtml(value);
</script>
