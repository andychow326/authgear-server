import { Controller } from "@hotwired/stimulus";

// Handle click link to submit form
// When clicking element with `data-submit-link`, it will perform click on
// element with `data-submit-form` that contains the same value
// e.g. data-submit-link="verify-identity-resend" and
//      data-submit-form="verify-identity-resend"
export class TransferClickController extends Controller {
  click(e: Event) {
    const link = this.element as HTMLElement;
    const buttonID = link.getAttribute("data-transfer-click-click");
    if (buttonID == null) {
      return;
    }

    const buttonTarget = document.querySelector(buttonID);

    if (buttonTarget instanceof HTMLElement) {
      e.preventDefault();
      buttonTarget.click();
    }
  }
}
