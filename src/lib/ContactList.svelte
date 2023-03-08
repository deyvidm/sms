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
            let i = recipientContacts.indexOf(contact)
            if (i > -1){
                recipientContacts.splice(i, 1)
            }
        }
        dispatch('message', {
            recipients: recipientContacts,
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
                        checked={false}
                        first={contact.first_name}
                        last={contact.last_name}
                        id={contact.id}
                    />
                {/each}
            </tbody>
        </table>
    </div>
</div>
