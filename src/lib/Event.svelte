<script lang="ts">
    import { beforeUpdate, createEventDispatcher, onMount } from 'svelte';
    import { detach } from 'svelte/internal';
    import EventAttendee from './EventAttendee.svelte';
    import {get50Attendees} from './pocketbase';
    import type { EventRecord, EventResponse, AttendeeResponse } from './pocketbase-types';

    export let event: EventResponse;
    export let active: boolean = false;
    let attendees = new Array<AttendeeResponse>();
    const dispatch = createEventDispatcher();

   async function loadAttendance() {
        // get50Attendees(event);
        await get50Attendees(event).then((result) => {
            attendees = result.items;
        });
    }

</script>

<div
    on:click|once={loadAttendance}
    class="mb-2 collapse border border-base-300 bg-base-100 rounded-box"
>
    <input type="checkbox" bind:checked={active} />
    <div class="collapse-title text-2xl font-medium">
        <span>{event.title}</span>
        <span class="absolute right-2">{event.start_date.slice(0, 10)}</span>
    </div>

    <div class="collapse-content">
        {#if active}
            <div class="mb-5">
                <button class="m-2 ml-0 btn btn-outline">Inite More</button>
                <button class="m-2 btn btn-outline">Follow Up</button>
                <button class="m-2 btn btn-outline">All Paid</button>
                <button class="m-2 btn btn-outline btn-error">Cancel</button>
            </div>
        {/if}

        {#each attendees as a}
            <EventAttendee who={a} />
        {/each}
    </div>
</div>
