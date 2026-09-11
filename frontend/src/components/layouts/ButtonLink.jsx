export default function ButtonLink({label, link, color, icon}) {
    const colorMap = {
        red: 'bg-red-500 hover:bg-red-600 text-white',
        blue: 'bg-blue-500 hover:bg-blue-600 text-white',
        green: 'bg-green-500 hover:bg-green-600 text-white' ,
        orange: 'bg-orange-500 hover:bg-orange-600 text-white',
        white: 'bg-white-500 hover:bg-gray-200'
    };
    function handlerClick(url) {
        window.open(url, '_blank');
    }
    return (
        <button className={`p-2 cursor-pointer rounded-2xl font-bold ${colorMap[color]} border-black-200 outline-2`} onClick={() => handlerClick(link)}>
            <div className={`flex gap-1`}>
                <img src={icon} alt={''} className="w-6 h-6 object-contain"/>
                <span>{label}</span>
            </div>
        </button>
    );
}