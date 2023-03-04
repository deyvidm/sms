<script lang="ts">
    import ContactRow from './ContactRow.svelte';
    import {  contacts } from './pocketbase';
    let yesall = false;
</script>

<div class="overflow-x-auto w-full">
    <table class="table w-full">
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
            {#each $contacts as contact}
                <ContactRow checked={false || yesall} first={contact.first_name} last={contact.last_name} id={contact.id} />
            {/each}
        </tbody>
    </table>
</div>

