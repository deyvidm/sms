<script lang="ts">
    import { onMount } from "svelte";
    import { userEvents } from "$lib/gin";
    import type { Event } from "$lib/Types";
    import EventRow from "$lib/EventRow.svelte"

    let events = new Array<Event>();

    onMount(async () => {
        console.log("mounted events page")
        if ($userEvents) {
            events = $userEvents;
        } 
        console.log(events)
    });

    let activeId = "";
    function handleMessage(event) {
        activeId = event.detail.eventId;
    }
</script>

<!-- <p>events: {events.length}</p> -->
{#each events as e}
     <EventRow event={e} active={activeId == e.id} on:message={handleMessage} />
{/each}
