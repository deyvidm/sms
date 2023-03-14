<script lang="ts">
    import { error } from '@sveltejs/kit';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import ContactList from './ContactList.svelte';
    import InputBox from './InputBox.svelte';
    import { createEvent, pb } from './pocketbase';
    import {
        EventStatusOptions,
        type AttendeeRecord,
        type ContactResponse,
        type EventRecord,
        AttendeeStatusOptions,
    } from './pocketbase-types';

    let formTitle: string;
    let content: string;
    let recipients = new Array<ContactResponse>();
    let submitButton: HTMLButtonElement;

    function handleMessage(event) {
        recipients = event.detail.recipients;
    }

    function create() {
        submitButton.classList.add("loading")
        // const d = new Date('05 October 2011 14:48 UTC');
        const d = new Date(Date.now());
        createEvent(<EventRecord>{
            organizer: pb.authStore.model?.id,
            title: formTitle,
            description: content,
            capacity: 666,
            start_date: d.toISOString(),
            end_date: d.toISOString(),
            send_invite_date: d.toISOString(),
            status: EventStatusOptions.active,
        }).then((eventRecord) => {
            recipients.forEach((r) => {
                pb.collection('attendee')
                    .create(<AttendeeRecord>{
                        event: eventRecord.id,
                        contact: r.id,
                        status: AttendeeStatusOptions['sending-invite'],
                        paid: false,
                    })
                    .catch((error) => {
                        console.log(error);
                    });
            });
        }).then(()=>{
            submitButton.classList.remove("loading")
            submitButton.classList.add("btn-disabled")
            submitButton.textContent = "Success"
        });
    }
</script>

<h2 class="mb-10 text-4xl font-extrabold dark:text-white">Create New Event</h2>

<div class="form-control w-full ">
    <div>
        <label class="label" for="event-title">
            <span class="text-xl font-bold text-xllabel-text">Title</span>
        </label>
        <input
            bind:value={formTitle}
            type="text"
            placeholder="Sunday Funday St Paddy's Vball"
            class="input input-bordered w-full max-w"
            id="event-title"
        />
    </div>
    <div class="mt-5">
        <label class="label" for="event-content">
            <span class="text-xl font-bold text-xllabel-text">SMS Content</span>
        </label>
        <textarea
            bind:value={content}
            class="w-full h-96 pb-10 textarea textarea-bordered"
            id="event-content"
        />
    </div>

    <!-- The button to open modal -->
    <label for="my-modal-5" class="mt-12 w-1/4 btn btn-outline btn-accent"
        >{recipients.length > 0 ? 'Edit Recipients' : 'Add Recipients'}</label
    >
    <input type="checkbox" id="my-modal-5" class="modal-toggle" />
    <div class="modal">
        <div class="modal-box w-11/12 max-w-5xl">
            <ContactList on:message={handleMessage} />
            <div class="modal-action">
                <label for="my-modal-5" class="btn">Finish</label>
            </div>
        </div>
    </div>
    <div class="mt-5" />
    {#if recipients.length > 0}
        <button bind:this={submitButton} on:click={create} class="mt-12 w-1/4 btn btn-active"
            >Create</button
        >
    {/if}
</div>
