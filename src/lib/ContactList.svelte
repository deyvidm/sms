<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import { contacts } from './pocketbase';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import type { ContactsResponse } from './pocketbase-types';
    import { createEventDispatcher } from 'svelte';
    
    let yesall = false;
    let remainingContacts = $contacts;
    let recipientContacts = new Array<ContactsResponse>();

    let contactIDMap = new Map<string, ContactsResponse>();
    $contacts.forEach((contact, i, parent) => {
        contactIDMap.set(contact.id, contact);
    });

    const dispatch = createEventDispatcher();
    onMount(() => {});
    afterUpdate(() => {});

    function handleMessage(event) {
        let id = event.detail.id;

        let contact = contactIDMap.get(id);
        if (contact == null) {
            return;
        }
        let checked = event.detail.checked;
        if (checked) {
            remainingContacts = remainingContacts.filter((c) => {
                c != contact;
            });
            recipientContacts.push(contact);
            recipientContacts = recipientContacts; //need this called to update the component

            // contact de-selected
        } else {
            recipientContacts = recipientContacts.filter((c) => {
                c != contact;
            });
            remainingContacts.push(contact);
            remainingContacts = remainingContacts; //need this called to update the component
            // https://svelte.dev/tutorial/updating-arrays-and-objects
        }

        dispatch('message', {
            recipients: recipientContacts,
            remaining: remainingContacts,
        });
    }
</script>

<div class="overflow-x-auto w-full">
    <div class="flex flex-col w-full lg:flex-row">
        <table class="table w-1/2">
            <thead>
                <tr>
                    <th>
                        <label>
                            <input bind:checked={yesall} type="checkbox" class="checkbox" />
                        </label>
                    </th>
                    <th>Name</th>
                </tr>
            </thead>

            <tbody>
                {#each remainingContacts as contact}
                    <ContactRow
                        on:message={handleMessage}
                        checked={false || yesall}
                        first={contact.first_name}
                        last={contact.last_name}
                        id={contact.id}
                    />
                {/each}
            </tbody>
        </table>
        <div class="divider lg:divider-horizontal" />
        <div class="flex flex-col w-full lg:flex-row">
            <!-- <div class=" w-5/12 p-5 shadow-xl"> -->
            <table class="table w-1/2">
                <thead>
                    <tr>
                        <th>
                            <label>
                                <input bind:checked={yesall} type="checkbox" class="checkbox" />
                            </label>
                        </th>
                        <th>Inbited</th>
                    </tr>
                </thead>

                <tbody>
                    {#each recipientContacts as contact}
                        <ContactRow
                            on:message={handleMessage}
                            checked={true || yesall}
                            first={contact.first_name}
                            last={contact.last_name}
                            id={contact.id}
                        />
                    {/each}
                </tbody>
            </table>
        </div>
    </div>
</div>
