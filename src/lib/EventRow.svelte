<script lang="ts">
    import { beforeUpdate, createEventDispatcher, onMount } from 'svelte';
    import EventAttendee from './EventAttendee.svelte';
    import { attendees, get50Attendees } from './pocketbase';
    import type { EventRecord, EventResponse } from './pocketbase-types';

    export let event: EventResponse;

    const dispatch = createEventDispatcher();

    function onClick() {
        get50Attendees(event);
        // console.log($attendees[0].expand.contact.first_name)
    }

    onMount(() => {
    });
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div on:click|once={onClick} class="mb-2 collapse border border-base-300 bg-base-100 rounded-box">
    <input type="checkbox" />
    <div class="collapse-title text-2xl font-medium">
        <span>{event.title}</span>
        <span class="absolute right-2">{event.start_date.slice(0,10)}</span>
    </div>
    <div class="collapse-content">
        {#each $attendees as a}
            <EventAttendee who={a}/>
        {/each}
    </div>
</div>
