import Item from './Item';

const ItemList = ({ courses, bandera }) => {
    if (bandera !== 1) { //esta bandera sirve para identificar si la lista de cursos viene del search o del boton mis cursos
        bandera = 0;
    }
    return (
        <section>
            <div>
                {courses.map((course) => (
                    <Item course={course} key={course.course_id} bandera={bandera} />
                ))}
            </div>
        </section>
    );
};

export default ItemList;