<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import { get50Contacts } from './pocketbase';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import type { AttendeeResponse, ContactResponse } from './pocketbase-types';
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    function handleMessage(event) {
        let id = event.detail.id;
        let checked = event.detail.checked;
        let contact = contactIDMap.get(id);

        if (contact == null) {
            return;
        }
        if (checked) {
            if (selectedContacts.indexOf(contact) < 0) {
                selectedContacts.push(contact);
            }
        } else {
            let i = selectedContacts.indexOf(contact);
            if (i > -1) {
                selectedContacts.splice(i, 1);
            }
        }
        selectedContacts = selectedContacts;
        remainingContacts = remainingContacts;
        dispatch('message', {
            recipients: selectedContacts,
        });
    }

    onMount(async () => {
        await get50Contacts().then((result) => {
            allContacts = result.items;
        });
        allContacts.forEach((contact, i, arr) => {
            contactIDMap.set(contact.id, contact);
        });

        ignore.forEach(id => {
            allContacts.forEach((contact,i, arr) =>{
                if (id == contact.id){
                    return
                }
                remainingContacts.push(contact)
            })
        });
        remainingContacts = remainingContacts;
    });

    // an array of ContactIDs to exclude when displaying the list
    // useful when trying to add more contacts to an existing event
    // i.e. hide attending contacts
    export let ignore = new Array<String>();

    let allContacts: Array<ContactResponse>;
    let remainingContacts = new Array<ContactResponse>();
    let selectedContacts = new Array<ContactResponse>();

    let contactIDMap = new Map<string, ContactResponse>();


</script>

<div class="overflow-x-auto w-full">
    <div class="flex flex-col w-full lg:flex-row">
        <table class="table w-full">
            <thead>
                <tr>
                    <th>
                        <label>
                            <!-- <input bind:checked={yesall} type="checkbox" class="checkbox" /> -->
                        </label>
                    </th>
                    <th>First</th>
                    <th>Last</th>
                    <th>Status</th>
                </tr>
            </thead>
            <tbody>
                {#each remainingContacts as contact}
                    <ContactRow on:message={handleMessage} checked={false} {contact} />
                {/each}
            </tbody>
        </table>
    </div>
</div>
