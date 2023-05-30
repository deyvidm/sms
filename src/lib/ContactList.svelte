<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';

    import { createEventDispatcher } from 'svelte';
    import type { Contact } from './gripes';

    export let contacts: Array<Contact>;

    const dispatch = createEventDispatcher();

    function handleMessage(event) {
        console.log(contacts.filter(c => c.id == event.detail.id))
        dispatch('message', {
            recipient: contacts.filter(c => c.id == event.detail.id)[0]
        });
    }

    beforeUpdate(() => {
    });

    onMount(async () => {
    });
    
</script>

<div class="overflow-x-auto w-full">
    <div class="flex flex-col w-full lg:flex-row">
        <table class="table w-full">
            <thead>
                <tr>
                    <th hidden={true} > 
                        <label>
                            <!-- <input bind:checked={yesall} type="checkbox" class="checkbox" /> -->
                        </label>
                    </th>
                    <th>First</th>
                    <th>Last</th>
                    <th>Phone</th>
                </tr>
            </thead>
            <tbody>
                {#each contacts as contact}
                    <ContactRow on:message={handleMessage} checked={false} {contact} />
                {/each}
            </tbody>
        </table>
    </div>
</div>
