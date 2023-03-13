<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import { contacts } from './pocketbase';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import type { ContactResponse } from './pocketbase-types';
    import { createEventDispatcher } from 'svelte';

    let yesall = false;
    let remainingContacts = $contacts;
    let recipientContacts = new Array<ContactResponse>();

    let contactIDMap = new Map<string, ContactResponse>();
    $contacts.forEach((contact, i, parent) => {
        contactIDMap.set(contact.id, contact);
    });

    const dispatch = createEventDispatcher();
    function handleMessage(event) {
        let id = event.detail.id;
        let checked = event.detail.checked;
        let contact = contactIDMap.get(id);

        if (contact == null) {
            return;
        }
        if (checked) {
            if (recipientContacts.indexOf(contact) < 0) {
                recipientContacts.push(contact);
            }
        } else {
            let i = recipientContacts.indexOf(contact);
            if (i > -1) {
                recipientContacts.splice(i, 1);
            }
        }
        recipientContacts = recipientContacts;
        remainingContacts = remainingContacts;
        dispatch('message', {
            recipients: recipientContacts,
        });
    }
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
                    <ContactRow
                        on:message={handleMessage}
                        checked={false}
                        contact={contact}
                    />
                {/each}
            </tbody>
        </table>
    </div>
</div>
