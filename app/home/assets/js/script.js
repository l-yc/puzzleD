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

window.onload = () => {
  let tbl = document.querySelector('#puzzles');

  specs = [
    { label: 'Name', init: tblInit.text },
    { label: 'Author', init: tblInit.text },
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
        tbody.appendChild(tr);
      });
      tbl.appendChild(tbody);
    })  
}

