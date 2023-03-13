<script lang="ts">
    import InputBox from './InputBox.svelte';
    import PhoneInputBox from './PhoneInputBox.svelte';
    import { pb, API, currentUser } from './pocketbase';
    import { ContactStatusOptions, type ContactRecord } from './pocketbase-types';

    // <input> tags at our disposal
    let firstname: InputBox;
    let lastname: InputBox;
    let phone: PhoneInputBox;
    let submit: HTMLButtonElement;
    let submitLabel: string = 'Create Contact';

    function createContact() {
        if (!verifyInputs() || !$currentUser) {
            return;
        }
        let data: ContactRecord = {
            first_name: firstname.Get(),
            last_name: lastname.Get(),
            phone: phone.Get(),
            owner: $currentUser.id,
            status: ContactStatusOptions.Pending,
        };
        API.createContact(data)
            .then((result) => setButtonStatus(true))
            .catch((error) => {
                setButtonStatus(false);
                console.log(error);
                submitLabel = 'shit';
            });
    }

    function setButtonStatus(success: boolean) {
        if (success) {
            submit.classList.add('btn-disabled', 'btn-success');
            submitLabel = 'Success!';
        } else {
            submit.classList.add('btn-disabled');
        }
    }
    function verifyInputs() {
        let results = [
            firstname.ValidateAndMark(),
            lastname.ValidateAndMark(),
            phone.ValidateAndMark(),
        ];
        return results.reduce((a, b) => a && b);
    }
</script>

<div class="form-control w-full max-w-xs">
    <InputBox bind:this={firstname} label="First Name" />
    <InputBox bind:this={lastname} label="Last Name" />
    <PhoneInputBox bind:this={phone} label="Phone" />
    <button bind:this={submit} on:click={createContact} class="mt-12 w-auto btn btn-active"
        >{submitLabel}</button
    >
</div>
