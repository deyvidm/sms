<script lang="ts">
    import { onMount } from "svelte";
    import type { Event } from "$lib/gripes";
    import EventRow from "$lib/EventRow.svelte";

    /** @type {import('./$types').PageData} */
    export let data;
    let events = new Array<Event>();

    onMount(async () => {
        console.log("mounted events page");
        if (data.events) {
            events = data.events;
        }
        console.log(events);
    });

    let activeId = "";
    function handleMessage(event) {
        activeId = event.detail.eventId;
    }
</script>

{#each events as e}
    <EventRow event={e} active={activeId == e.id} on:message={handleMessage} />
{/each}
