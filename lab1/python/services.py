import openpyxl
from datetime import datetime, date


def get_complex_field(dt, f_name):
    names = f_name.split('.')
    for name in names:
        dt = dt[name]
    return dt


def import_to_excel(json: dict):
    wb = openpyxl.Workbook()
    wb.remove(wb.active)
    for field in json['fields']:
        ws = wb.create_sheet(field['title'])
        fields = {i['field_name']: index for index, i in enumerate(field['entity_fields'])}
        ws.append([i['title'] for i in field['entity_fields']])

        for row in json['data'][field['field_name']]:
            data = [None for _ in range(len(field['entity_fields']))]
            for entity_field in field['entity_fields']:
                data[fields[entity_field['field_name']]] = get_complex_field(row, entity_field['field_name'])
            ws.append(data)

    wb.save(f'{str(datetime.now()).replace(":", "-")}.xlsx')


if __name__ == '__main__':
    import_to_excel(
        {'faculties': [],
         'specializations': [],
         'teachers': [],
         'courses': [],
         'emails': [],
         'students': []})
