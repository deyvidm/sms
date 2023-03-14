<script lang="ts">
    import { beforeUpdate, createEventDispatcher, onMount } from 'svelte';
    import { detach } from 'svelte/internal';
    import EventAttendee from './EventAttendee.svelte';
    import { get50Attendees } from './pocketbase';
    import ContactList from './ContactList.svelte';
    import {
        type EventRecord,
        type EventResponse,
        type AttendeeResponse,
        AttendeeStatusOptions,
    } from './pocketbase-types';

    const dispatch = createEventDispatcher();

    async function loadAttendance() {
        await get50Attendees(event).then((result) => {
            attendees = result.items;
            attendees.forEach((a) => {
                if (a.status == AttendeeStatusOptions.accepted) {
                    attendeesConfirmed++;
                }
                ignoreList.push(a.contact)
            });
            ignoreList = ignoreList;
        });
        console.log(ignoreList)
    }

    function handleMessage(event) {
        console.log(event.detail);
    }


    export let event: EventResponse;
    export let active: boolean = false;
    let attendees = new Array<AttendeeResponse>();
    let attendeesConfirmed = 0;
    let ignoreList = new Array<string>();
</script>

<div
    on:click|once={loadAttendance}
    class="mb-5 collapse border border-base-300 bg-base-200 rounded-box {active
        ? 'bg-base-300'
        : ''}"
>
    <input type="checkbox" bind:checked={active} />
    <div class="collapse-title text-2xl font-medium">
        <span>{event.title}</span>
        <span class="absolute right-2">{event.start_date.slice(0, 10)}</span>
    </div>

    <div class="collapse-content">
        {#if active}
            <div class="mb-5">
                <div class="mb-2">attendance: {attendeesConfirmed}/{attendees.length}</div>
                <label for="invite-more-modal" class="m-2 ml-0 btn btn-outline">Invite More</label>
                <input type="checkbox" id="invite-more-modal" class="modal-toggle" />
                <div class="modal">
                    <div class="modal-box w-11/12 max-w-5xl">
                        <ContactList on:message={handleMessage} ignore={ignoreList} />
                        <div class="modal-action">
                            <label for="invite-more-modal" class="btn">Finish</label>
                        </div>
                    </div>
                </div>

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
