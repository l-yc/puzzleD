console.clear();

tblInit = {
  text: function(field, data) {
    let td = document.createElement('td');
    td.innerHTML = data;
    return td;
  },
  download: function(field, data) {
    let td = document.createElement('td');
    let btn = document.createElement('button');
    btn.innerHTML = 'Download';
    td.appendChild(btn);
    return td;
  },
  submit: function(field, data) {
    let td = document.createElement('td');
    let input = document.createElement('input');
    td.appendChild(input);
    return td;
  }
}

class Modal {
  constructor() {
    let container = document.createElement('div');
    container.classList.add('modal');

    let wrapper = document.createElement('div');
    wrapper.classList.add('content-wrapper');
    container.appendChild(wrapper);

    let closeBtn = document.createElement('span');
    closeBtn.classList.add('close-button');
    closeBtn.innerHTML = '&times;';
    wrapper.appendChild(closeBtn);

    let header = document.createElement('div');
    header.classList.add('header');
    wrapper.appendChild(header);

    let overflowWrapper = document.createElement('div');
    wrapper.appendChild(overflowWrapper);

    let content = document.createElement('div');
    content.classList.add('content');
    overflowWrapper.appendChild(content);

    let footer = document.createElement('div');
    footer.classList.add('footer');
    wrapper.appendChild(footer);

    this.el = container;
    this.header = header;
    this.content = content;
    this.footer = footer;
    this.closeBtn = closeBtn;

    this.closeBtn.addEventListener('click', e => {
      this.toggle();
    });
    this.el.addEventListener('click', e => {
      console.log(e.target, this.el);
      if (e.target == this.el) this.toggle();
    });
  }

  setHeaderHTML(htmlContent) {
    this.header.innerHTML = htmlContent;
  }

  setContentHTML(htmlContent) {
    this.content.innerHTML = htmlContent;
  }

  setFooterHTML(htmlContent) {
    console.log(this.footer);
    this.footer.innerHTML = htmlContent;
  }

  setHeader(...content) {
    this.header.innerHTML = '';
    this.header.appendChild(...content);
  }

  setContent(...content) {
    this.content.innerHTML = '';
    this.content.appendChild(...content);
  }

  setFooter(...content) {
    this.footer.innerHTML = '';
    this.footer.appendChild(...content);
  }

  toggle() {
    this.el.classList.toggle('show');
  }
};
modal = new Modal();

function displayPuzzle(puzzle) {
  let header = `
    <h1>${puzzle.Name}</h1>
    <h5>Authors: <a href="#">${puzzle.Authors}</a></h5>
    <hr>
  `;
  modal.setHeaderHTML(header);

  modal.setContentHTML(puzzle.Description);
  modal.toggle();
}

window.onload = () => {
  // set up flag submission modal
  document.body.appendChild(modal.el);

  let footer = document.createElement('form');
  footer.id = 'submit-flag';

  let flagInput = document.createElement('input');
  flagInput.type = 'text';
  flagInput.placeholder = 'flag'; // TODO: change this to go with dynamic flag format!
  footer.appendChild(flagInput);

  let submit = document.createElement('div');
  submit.classList.add('button');
  submit.innerHTML = '<i class="las la-arrow-circle-right"></i>';
  submit.addEventListener('click', e => {
    alert('submit flag: ' + flagInput.value); // TODO: call the api! 
  });
  footer.appendChild(submit);

  modal.setFooter(footer);

  // set up puzzles table
  let tbl = document.querySelector('#puzzles');

  specs = [
    { label: 'Name', init: tblInit.text },
    { label: 'Authors', init: tblInit.text },
    { label: 'Score', init: tblInit.text },
    { label: 'Attachments', init: tblInit.download },
    { label: 'Submit', init: tblInit.submit },
  ]

  // create table
  let thead = document.createElement('thead');
  let tr = document.createElement('tr');
  specs.forEach(field => {
    let th = document.createElement('th');
    th.innerHTML = field.label;
    tr.appendChild(th);
  })
  thead.appendChild(tr);
  tbl.appendChild(thead);

  // load data
  fetch('/api/puzzles.json', { method: 'GET' })
    .then(res => res.json())
    .then(json => {
      console.log('fetched', json);

      let tbody = document.createElement('tbody');
      json.forEach(entry => {
        let tr = document.createElement('tr');
        specs.forEach(field => {
          tr.appendChild(field.init(field, entry[field.label]));
        })
        tr.addEventListener('click', e => {
          if (e.target.matches('td') || e.target == tr) displayPuzzle(entry);
        });
        tbody.appendChild(tr);
      });
      tbl.appendChild(tbody);
    })  
}

