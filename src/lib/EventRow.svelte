<script lang="ts">
    import { beforeUpdate, createEventDispatcher, onMount } from 'svelte';
    import { detach } from 'svelte/internal';
    import EventAttendee from './EventAttendee.svelte';
    import ContactList from './ContactList.svelte';
    import type { Event } from './gripes';
    import { apiClient } from './gin';

    const dispatch = createEventDispatcher();

    export let event: Event;
    export let active: boolean = false;
    let attendees = new Array();
    let attendeesConfirmed = 0;
    let ignoreList = new Array<string>();

    async function loadAttendance() {
        await apiClient.GetInvitations(event.id).then((resp) => {
            // could also scan resp.status == success
            attendees = resp.data.invites;
            attendees.forEach((a) => {
                if (a.status == 'accepted') {
                    attendeesConfirmed++;
                }
                //     ignoreList.push(a.contact)
            });
            // ignoreList = ignoreList;
            // return resp.data;
        });
        // await apiClient.GetAttendees(event).then((result) => {
        //     attendees = result.items;
        //     attendees.forEach((a) => {
        //         if (a.status == AttendeeStatusOptions.accepted) {
        //             attendeesConfirmed++;
        //         }
        //         ignoreList.push(a.contact)
        //     });
        //     ignoreList = ignoreList;
        // });
    }

    function handleMessage(event) {
        console.log(event.detail);
    }
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
                <!-- <div class="mb-2">attendance: {attendeesConfirmed}/{attendees.length}</div> -->
                <label for="invite-more-modal" class="m-2 ml-0 btn btn-outline">Invite More</label>
                <input type="checkbox" id="invite-more-modal" class="modal-toggle" />
                <div class="modal">
                    <div class="modal-box w-11/12 max-w-5xl">
                        <ContactList contacts={ignoreList} on:message={handleMessage} />
                        <div class="modal-action">
                            <label for="invite-more-modal" class="btn">Invite</label>
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
